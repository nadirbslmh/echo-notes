package database

import (
	"echo-notes/model"
	"echo-notes/util"
	"errors"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

var (
	DB_USERNAME string = util.GetConfig("DB_USERNAME")
	DB_PASSWORD string = util.GetConfig("DB_PASSWORD")
	DB_NAME     string = util.GetConfig("DB_NAME")
	DB_HOST     string = util.GetConfig("DB_HOST")
	DB_PORT     string = util.GetConfig("DB_PORT")
)

func Connect() {
	var err error

	var dsn string = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		DB_USERNAME,
		DB_PASSWORD,
		DB_HOST,
		DB_PORT,
		DB_NAME,
	)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("error when connecting to the database: %s", err)
	}

	log.Println("connected to the database")

	DB.AutoMigrate(&model.Note{}, &model.Category{}, &model.User{})
}

func InitTestDB() {
	var err error

	var dsn string = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		DB_USERNAME,
		DB_PASSWORD,
		DB_HOST,
		DB_PORT,
		util.GetConfig("DB_TEST_NAME"),
	)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("error when connecting to the database: %s", err)
	}

	log.Println("connected to the database")

	DB.AutoMigrate(&model.Note{}, &model.Category{}, &model.User{})
}

func SeedCategory() model.Category {
	var category model.Category = model.Category{
		Name: "sample",
	}

	if err := DB.Create(&category).Error; err != nil {
		panic(err)
	}

	var createdCategory model.Category

	DB.Last(&createdCategory)

	return createdCategory
}

func SeedNote() model.Note {

	category := SeedCategory()

	var note model.Note = model.Note{
		Title:      "test",
		Content:    "test",
		CategoryID: category.ID,
	}

	if err := DB.Create(&note).Error; err != nil {
		panic(err)
	}

	var createdNote model.Note

	DB.Last(&createdNote)

	return createdNote
}

func SeedUser() model.User {
	password, _ := bcrypt.GenerateFromPassword([]byte("123123"), bcrypt.DefaultCost)

	var user model.User = model.User{
		Email:    "testing@mail.com",
		Password: string(password),
	}

	if err := DB.Create(&user).Error; err != nil {
		panic(err)
	}

	var createdUser model.User

	DB.Last(&createdUser)

	createdUser.Password = "123123"

	return createdUser
}

func CleanSeeders() {
	categoryResult := DB.Exec("DELETE FROM categories")
	itemResult := DB.Exec("DELETE FROM notes")
	userResult := DB.Exec("DELETE FROM users")

	var isFailed bool = itemResult.Error != nil || userResult.Error != nil || categoryResult.Error != nil

	if isFailed {
		panic(errors.New("error when cleaning up seeders"))
	}

	log.Println("Seeders are cleaned up successfully")
}
