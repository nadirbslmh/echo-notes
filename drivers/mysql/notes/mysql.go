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

	nr.conn.Find(&rec)

	noteDomain := []notes.Domain{}

	for _, note := range rec {
		noteDomain = append(noteDomain, note.ToDomain())
	}

	return noteDomain
}

func (nr *noteRepository) GetByID(id string) notes.Domain {
	var note Note

	nr.conn.First(&note, "id = ?", id)

	return note.ToDomain()
}

func (nr *noteRepository) Create(noteDomain *notes.Domain) notes.Domain {
	var createdNote notes.Domain

	result := nr.conn.Create(&noteDomain)

	result.Last(&createdNote)

	return createdNote
}

func (nr *noteRepository) Update(id string, noteDomain *notes.Domain) notes.Domain {
	var note notes.Domain = nr.GetByID(id)

	note.Title = noteDomain.Title
	note.Content = noteDomain.Content
	note.CategoryID = noteDomain.CategoryID

	nr.conn.Save(&note)

	return note
}

func (nr *noteRepository) Delete(id string) bool {
	var note notes.Domain = nr.GetByID(id)

	result := nr.conn.Delete(&note)

	if result.RowsAffected == 0 {
		return false
	}

	return true
}

func (nr *noteRepository) Restore(id string) notes.Domain {
	var trashedNote notes.Domain

	nr.conn.Unscoped().First(&trashedNote, "id = ?", id)

	trashedNote.DeletedAt = gorm.DeletedAt{}

	nr.conn.Unscoped().Save(&trashedNote)

	return trashedNote
}

func (nr *noteRepository) ForceDelete(id string) bool {
	var note notes.Domain = nr.GetByID(id)

	result := nr.conn.Unscoped().Delete(&note)

	if result.RowsAffected == 0 {
		return false
	}

	return true
}
