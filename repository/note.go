package repository

import (
	"echo-notes/database"
	"echo-notes/model"
)

type NoteRepositoryImpl struct{}

func (n *NoteRepositoryImpl) GetAll() []model.Note {
	var notes []model.Note

	database.DB.Find(&notes)

	return notes
}

func (c *NoteRepositoryImpl) GetByID(id string) model.Note {
	var note model.Note

	database.DB.First(&note, "id = ?", id)

	return note
}

func (n *NoteRepositoryImpl) Create(input model.NoteInput) model.Note {
	var newNote model.Note = model.Note{
		Title:   input.Title,
		Content: input.Content,
	}

	var createdNote model.Note = model.Note{}

	result := database.DB.Create(&newNote)

	result.Last(&createdNote)

	return createdNote
}

func (n *NoteRepositoryImpl) Update(id string, input model.NoteInput) model.Note {
	var note model.Note = n.GetByID(id)

	note.Title = input.Title
	note.Content = input.Content

	database.DB.Save(&note)

	return note
}

func (n *NoteRepositoryImpl) Delete(id string) bool {
	var note model.Note = n.GetByID(id)

	result := database.DB.Delete(&note)

	if result.RowsAffected == 0 {
		return false
	}

	return true
}
