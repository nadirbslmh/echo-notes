package mysql_driver

import (
	"echo-notes/drivers/mysql/categories"
	"echo-notes/drivers/mysql/notes"
	"echo-notes/drivers/mysql/users"
	"echo-notes/model"
	"errors"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ConfigDB struct {
	DB_USERNAME string
	DB_PASSWORD string
	DB_NAME     string
	DB_HOST     string
	DB_PORT     string
}

func (config *ConfigDB) InitDB() *gorm.DB {
	var err error

	var dsn string = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.DB_USERNAME,
		config.DB_PASSWORD,
		config.DB_HOST,
		config.DB_PORT,
		config.DB_NAME,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("error when connecting to the database: %s", err)
	}

	log.Println("connected to the database")

	return db
}

func DBMigrate(db *gorm.DB) {
	db.AutoMigrate(&notes.Note{}, &categories.Category{}, &users.User{})
}

func SeedCategory(db *gorm.DB) model.Category {
	var category model.Category = model.Category{
		Name: "sample",
	}

	if err := db.Create(&category).Error; err != nil {
		panic(err)
	}

	var createdCategory model.Category

	db.Last(&createdCategory)

	return createdCategory
}

func SeedNote(db *gorm.DB) model.Note {

	category := SeedCategory(db)

	var note model.Note = model.Note{
		Title:      "test",
		Content:    "test",
		CategoryID: category.ID,
	}

	if err := db.Create(&note).Error; err != nil {
		panic(err)
	}

	var createdNote model.Note

	db.Last(&createdNote)

	return createdNote
}

func SeedUser(db *gorm.DB) model.User {
	password, _ := bcrypt.GenerateFromPassword([]byte("123123"), bcrypt.DefaultCost)

	var user model.User = model.User{
		Email:    "testing@mail.com",
		Password: string(password),
	}

	if err := db.Create(&user).Error; err != nil {
		panic(err)
	}

	var createdUser model.User

	db.Last(&createdUser)

	createdUser.Password = "123123"

	return createdUser
}

func CleanSeeders(db *gorm.DB) {
	db.Exec("SET FOREIGN_KEY_CHECKS = 0")

	categoryResult := db.Exec("DELETE FROM categories")
	itemResult := db.Exec("DELETE FROM notes")
	userResult := db.Exec("DELETE FROM users")

	var isFailed bool = itemResult.Error != nil || userResult.Error != nil || categoryResult.Error != nil

	if isFailed {
		panic(errors.New("error when cleaning up seeders"))
	}

	log.Println("Seeders are cleaned up successfully")
}
