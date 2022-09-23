package service

import (
	"echo-notes/model"
	"echo-notes/repository"
)

type NoteService struct {
	Repository repository.NoteRepository
}

func New() NoteService {
	return NoteService{
		Repository: &repository.NoteRepositoryImpl{},
	}
}

func (n *NoteService) GetAll() []model.Note {
	return n.Repository.GetAll()
}

func (n *NoteService) GetByID(id string) model.Note {
	return n.Repository.GetByID(id)
}

func (n *NoteService) Create(input model.NoteInput) model.Note {
	return n.Repository.Create(input)
}

func (n *NoteService) Update(id string, input model.NoteInput) model.Note {
	return n.Repository.Update(id, input)
}

func (n *NoteService) Delete(id string) bool {
	return n.Repository.Delete(id)
}
