package repository

import (
	"echo-notes/auth"
	"echo-notes/database"
	"echo-notes/model"

	"golang.org/x/crypto/bcrypt"
)

type AuthRepositoryImpl struct{}

func (a *AuthRepositoryImpl) Register(input model.UserInput) model.User {
	password, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)

	var newUser model.User = model.User{
		Email:    input.Email,
		Password: string(password),
	}

	var createdUser model.User = model.User{}

	result := database.DB.Create(&newUser)

	result.Last(&createdUser)

	return createdUser
}

func (a *AuthRepositoryImpl) Login(input model.UserInput) string {
	var user model.User = model.User{}

	database.DB.First(&user, "email = ?", input.Email)

	if user.ID == 0 {
		return ""
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))

	if err != nil {
		return ""
	}

	token := auth.CreateToken(user.ID)

	return token
}
