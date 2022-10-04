package notes

import (
	"echo-notes/businesses/notes"

	"gorm.io/gorm"
)

type noteRepository struct {
	conn *gorm.DB
}

func NewMySQLRepository(conn *gorm.DB) notes.Repository {
	return &noteRepository{
		conn: conn,
	}
}

func (nr *noteRepository) GetAll() []notes.Domain {
	var rec []Note

	nr.conn.Preload("Category").Find(&rec)

	noteDomain := []notes.Domain{}

	for _, note := range rec {
		noteDomain = append(noteDomain, note.ToDomain())
	}

	return noteDomain
}

func (nr *noteRepository) GetByID(id string) notes.Domain {
	var note Note

	nr.conn.Preload("Category").First(&note, "id = ?", id)

	return note.ToDomain()
}

func (nr *noteRepository) Create(noteDomain *notes.Domain) notes.Domain {
	rec := FromDomain(noteDomain)

	result := nr.conn.Create(&rec)

	result.Last(&rec)

	return rec.ToDomain()
}

func (nr *noteRepository) Update(id string, noteDomain *notes.Domain) notes.Domain {
	var note notes.Domain = nr.GetByID(id)

	updatedNote := FromDomain(&note)

	updatedNote.Title = noteDomain.Title
	updatedNote.Content = noteDomain.Content
	updatedNote.CategoryID = noteDomain.CategoryID

	nr.conn.Save(&updatedNote)

	return updatedNote.ToDomain()
}

func (nr *noteRepository) Delete(id string) bool {
	var note notes.Domain = nr.GetByID(id)

	deletedNote := FromDomain(&note)

	result := nr.conn.Delete(&deletedNote)

	if result.RowsAffected == 0 {
		return false
	}

	return true
}

func (nr *noteRepository) Restore(id string) notes.Domain {
	var trashedNote notes.Domain

	trashed := FromDomain(&trashedNote)

	nr.conn.Unscoped().First(&trashed, "id = ?", id)

	trashed.DeletedAt = gorm.DeletedAt{}

	nr.conn.Unscoped().Save(&trashed)

	return trashed.ToDomain()
}

func (nr *noteRepository) ForceDelete(id string) bool {
	var note notes.Domain = nr.GetByID(id)

	deletedNote := FromDomain(&note)

	result := nr.conn.Unscoped().Delete(&deletedNote)

	if result.RowsAffected == 0 {
		return false
	}

	return true
}
