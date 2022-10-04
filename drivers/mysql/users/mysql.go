package users

import (
	"echo-notes/auth"
	"echo-notes/businesses/users"
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type userRepository struct {
	conn *gorm.DB
}

func NewMySQLRepository(conn *gorm.DB) users.Repository {
	return &userRepository{
		conn: conn,
	}
}

func (ur *userRepository) Register(userDomain *users.Domain) users.Domain {
	password, _ := bcrypt.GenerateFromPassword([]byte(userDomain.Password), bcrypt.DefaultCost)

	rec := FromDomain(userDomain)

	rec.Password = string(password)

	// var createdUser users.Domain

	// createdUser.Password = string(password)

	result := ur.conn.Create(&rec)

	result.Last(&rec)

	return rec.ToDomain()
}

func (ur *userRepository) Login(userDomain *users.Domain) string {
	var user users.Domain

	ur.conn.First(&user, "email = ?", userDomain.Email)

	if user.ID == 0 {
		fmt.Println("user not found")
		return ""
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userDomain.Password))

	if err != nil {
		fmt.Println("password failed!")
		return ""
	}

	token := auth.CreateToken(user.ID)

	return token
}
