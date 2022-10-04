package drivers

import (
	categoryDomain "echo-notes/businesses/categories"
	categoryDB "echo-notes/drivers/mysql/categories"

	noteDomain "echo-notes/businesses/notes"
	noteDB "echo-notes/drivers/mysql/notes"

	userDomain "echo-notes/businesses/users"
	userDB "echo-notes/drivers/mysql/users"

	"gorm.io/gorm"
)

func NewCategoryRepository(conn *gorm.DB) categoryDomain.Repository {
	return categoryDB.NewMySQLRepository(conn)
}

func NewNoteRepository(conn *gorm.DB) noteDomain.Repository {
	return noteDB.NewMySQLRepository(conn)
}

func NewUserRepository(conn *gorm.DB) userDomain.Repository {
	return userDB.NewMySQLRepository(conn)
}
