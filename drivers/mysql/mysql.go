package mysql_driver

import (
	"echo-notes/drivers/mysql/categories"
	"echo-notes/drivers/mysql/notes"
	"echo-notes/drivers/mysql/users"
	"echo-notes/util"
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

func SeedCategory(db *gorm.DB) categories.Category {
	category, _ := util.CreateFaker[categories.Category]()

	if err := db.Create(&category).Error; err != nil {
		panic(err)
	}

	db.Last(&category)

	return category
}

func SeedNote(db *gorm.DB) notes.Note {

	category := SeedCategory(db)

	note, _ := util.CreateFaker[notes.Note]()

	note.CategoryID = category.ID

	if err := db.Create(&note).Error; err != nil {
		panic(err)
	}

	db.Last(&note)

	return note
}

func SeedUser(db *gorm.DB) users.User {
	password, _ := bcrypt.GenerateFromPassword([]byte("123123"), bcrypt.DefaultCost)

	fakeUser, _ := util.CreateFaker[users.User]()

	userRecord := users.User{
		Email:    fakeUser.Email,
		Password: string(password),
	}

	if err := db.Create(&userRecord).Error; err != nil {
		panic(err)
	}

	var foundUser users.User

	db.Last(&foundUser)

	foundUser.Password = "123123"

	return foundUser
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
