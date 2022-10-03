package service

import (
	"echo-notes/model"
	"echo-notes/repository"
	userRepository "echo-notes/repository/users"
)

type AuthService struct {
	Repository repository.AuthRepository
}

func NewAuthService() AuthService {
	return AuthService{
		Repository: &userRepository.AuthRepositoryImpl{},
	}
}

func (a *AuthService) Register(input model.UserInput) model.User {
	return a.Repository.Register(input)
}

func (a *AuthService) Login(input model.UserInput) string {
	return a.Repository.Login(input)
}
