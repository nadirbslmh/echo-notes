package repository

import (
	"echo-notes/database"
	"echo-notes/model"
	"log"
	"strconv"
	"time"
)

type NoteRepositoryImpl struct{}

func (n *NoteRepositoryImpl) GetAll() []model.Note {
	rows, err := database.DB.Query("SELECT * FROM notes ORDER BY created_at")

	if err != nil {
		log.Fatalf("error when fetching data: %s", err)
	}

	note := model.Note{}

	notes := []model.Note{}

	for rows.Next() {
		var id int
		var title, content string
		var createdAt time.Time

		err = rows.Scan(&id, &title, &content, &createdAt)

		if err != nil {
			log.Fatalf("error when fetching data: %s", err)
		}

		note.ID = id
		note.Title = title
		note.Content = content
		note.CreatedAt = createdAt

		notes = append(notes, note)
	}

	return notes
}

func (c *NoteRepositoryImpl) GetByID(id string) model.Note {
	rows, err := database.DB.Query("SELECT * FROM notes WHERE id=?", id)

	if err != nil {
		log.Fatalf("error when fetching data: %s", err)
	}

	foundNote := model.Note{}

	for rows.Next() {
		var id int
		var title, content string
		var createdAt time.Time

		err = rows.Scan(&id, &title, &content, &createdAt)
		if err != nil {
			log.Fatalf("error when fetching data: %s", err)
		}

		foundNote.ID = id
		foundNote.Title = title
		foundNote.Content = content
		foundNote.CreatedAt = createdAt
	}

	return foundNote
}

func (n *NoteRepositoryImpl) Create(input model.NoteInput) model.Note {
	statement, err := database.DB.Prepare("INSERT INTO notes(title,content) VALUES(?,?)")

	if err != nil {
		log.Fatalf("error when adding data: %s", err)
	}

	result, err := statement.Exec(input.Title, input.Content)

	if err != nil {
		log.Fatalf("error when adding data: %s", err)
	}

	insertedId, err := result.LastInsertId()

	if err != nil {
		log.Fatalf("error when adding data: %s", err)
	}

	rows, err := database.DB.Query("SELECT * FROM notes WHERE id=?", int(insertedId))

	if err != nil {
		log.Fatalf("error when fetching data: %s", err)
	}

	createdNote := model.Note{}

	for rows.Next() {
		var id int
		var title, content string
		var createdAt time.Time

		err = rows.Scan(&id, &title, &content, &createdAt)
		if err != nil {
			log.Fatalf("error when fetching data: %s", err)
		}

		createdNote.ID = id
		createdNote.Title = title
		createdNote.Content = content
		createdNote.CreatedAt = createdAt
	}

	return createdNote
}

func (n *NoteRepositoryImpl) Update(id string, input model.NoteInput) model.Note {
	statement, err := database.DB.Prepare("UPDATE notes SET title=?, content=? WHERE id=?")

	if err != nil {
		log.Fatalf("error when updating data: %s", err)
	}

	_, err = statement.Exec(input.Title, input.Content, id)

	if err != nil {
		log.Fatalf("error when updating data: %s", err)
	}

	noteId, _ := strconv.Atoi(id)

	rows, err := database.DB.Query("SELECT * FROM notes WHERE id=?", noteId)

	if err != nil {
		log.Fatalf("error when fetching data: %s", err)
	}

	updatedNote := model.Note{}

	for rows.Next() {
		var id int
		var title, content string
		var createdAt time.Time

		err = rows.Scan(&id, &title, &content, &createdAt)
		if err != nil {
			log.Fatalf("error when fetching data: %s", err)
		}

		updatedNote.ID = id
		updatedNote.Title = title
		updatedNote.Content = content
		updatedNote.CreatedAt = createdAt
	}

	return updatedNote
}

func (n *NoteRepositoryImpl) Delete(id string) bool {
	statement, err := database.DB.Prepare("DELETE FROM notes WHERE id=?")

	if err != nil {
		log.Fatalf("error when deleting data: %s", err)
	}

	result, err := statement.Exec(id)

	if err != nil {
		log.Fatalf("error when deleting data: %s", err)
	}

	rowsAffected, _ := result.RowsAffected()

	if rowsAffected == 0 {
		return false
	}

	return true
}
