package main

import (
	_driverFactory "echo-notes/drivers"
	"echo-notes/util"
	"log"

	_noteUseCase "echo-notes/businesses/notes"
	_noteController "echo-notes/controller/notes"

	_categoryUseCase "echo-notes/businesses/categories"
	_categoryController "echo-notes/controller/categories"

	_userUseCase "echo-notes/businesses/users"
	_userController "echo-notes/controller/users"

	_dbDriver "echo-notes/drivers/mysql"

	_middleware "echo-notes/app/middlewares"
	_routes "echo-notes/app/routes"

	echo "github.com/labstack/echo/v4"
)

func main() {
	configDB := _dbDriver.ConfigDB{
		DB_USERNAME: util.GetConfig("DB_USERNAME"),
		DB_PASSWORD: util.GetConfig("DB_PASSWORD"),
		DB_HOST:     util.GetConfig("DB_HOST"),
		DB_PORT:     util.GetConfig("DB_PORT"),
		DB_NAME:     util.GetConfig("DB_NAME"),
	}

	db := configDB.InitDB()

	_dbDriver.DBMigrate(db)

	configJWT := _middleware.ConfigJWT{
		SecretJWT:       util.GetConfig("JWT_SECRET_KEY"),
		ExpiresDuration: 1,
	}

	e := echo.New()

	categoryRepo := _driverFactory.NewCategoryRepository(db)
	categoryUsecase := _categoryUseCase.NewCategoryUsecase(categoryRepo)
	categoryCtrl := _categoryController.NewCategoryController(categoryUsecase)

	noteRepo := _driverFactory.NewNoteRepository(db)
	noteUsecase := _noteUseCase.NewNoteUsecase(noteRepo)
	noteCtrl := _noteController.NewNoteController(noteUsecase)

	userRepo := _driverFactory.NewUserRepository(db)
	userUsecase := _userUseCase.NewUserUsecase(userRepo)
	userCtrl := _userController.NewAuthController(userUsecase)

	routesInit := _routes.ControllerList{
		JWTMiddleware:      configJWT.Init(),
		CategoryController: *categoryCtrl,
		NoteController:     *noteCtrl,
		AuthController:     *userCtrl,
	}

	routesInit.RouteRegister(e)

	log.Fatal(e.Start(":1323"))
}
