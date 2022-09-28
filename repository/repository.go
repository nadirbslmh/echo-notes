package repository

import "echo-notes/model"

type NoteRepository interface {
	GetAll() []model.Note
	GetByID(id string) model.Note
	Create(input model.NoteInput) model.Note
	Update(id string, input model.NoteInput) model.Note
	Delete(id string) bool
	Restore(id string) model.Note
	ForceDelete(id string) bool
}

type AuthRepository interface {
	Register(input model.UserInput) model.User
	Login(input model.UserInput) string
}
