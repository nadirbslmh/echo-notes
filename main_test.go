package main

import (
	_driverFactory "echo-notes/drivers"
	"echo-notes/util"
	"encoding/json"
	"net/http"
	"strconv"
	"testing"

	_noteUseCase "echo-notes/businesses/notes"
	_noteController "echo-notes/controller/notes"
	"echo-notes/controller/users/request"

	_categoryUseCase "echo-notes/businesses/categories"
	_categoryController "echo-notes/controller/categories"

	_userUseCase "echo-notes/businesses/users"
	_userController "echo-notes/controller/users"

	_dbDriver "echo-notes/drivers/mysql"
	"echo-notes/drivers/mysql/categories"
	"echo-notes/drivers/mysql/notes"
	"echo-notes/drivers/mysql/users"

	_middleware "echo-notes/app/middlewares"
	_routes "echo-notes/app/routes"

	echo "github.com/labstack/echo/v4"
	"github.com/steinfletcher/apitest"
)

func newApp() *echo.Echo {
	configDB := _dbDriver.ConfigDB{
		DB_USERNAME: util.GetConfig("DB_USERNAME"),
		DB_PASSWORD: util.GetConfig("DB_PASSWORD"),
		DB_HOST:     util.GetConfig("DB_HOST"),
		DB_PORT:     util.GetConfig("DB_PORT"),
		DB_NAME:     util.GetConfig("DB_TEST_NAME"),
	}

	db := configDB.InitDB()

	_dbDriver.DBMigrate(db)

	configJWT := _middleware.ConfigJWT{
		SecretJWT:       util.GetConfig("JWT_SECRET_KEY"),
		ExpiresDuration: 1,
	}

	configLogger := _middleware.ConfigLogger{
		Format: "[${time_rfc3339}] ${status} ${method} ${host} ${path} ${latency_human}" + "\n",
	}

	e := echo.New()

	categoryRepo := _driverFactory.NewCategoryRepository(db)
	categoryUsecase := _categoryUseCase.NewCategoryUsecase(categoryRepo)
	categoryCtrl := _categoryController.NewCategoryController(categoryUsecase)

	noteRepo := _driverFactory.NewNoteRepository(db)
	noteUsecase := _noteUseCase.NewNoteUsecase(noteRepo)
	noteCtrl := _noteController.NewNoteController(noteUsecase)

	userRepo := _driverFactory.NewUserRepository(db)
	userUsecase := _userUseCase.NewUserUsecase(userRepo, &configJWT)
	userCtrl := _userController.NewAuthController(userUsecase)

	routesInit := _routes.ControllerList{
		LoggerMiddleware:   configLogger.Init(),
		JWTMiddleware:      configJWT.Init(),
		CategoryController: *categoryCtrl,
		NoteController:     *noteCtrl,
		AuthController:     *userCtrl,
	}

	routesInit.RouteRegister(e)

	return e
}

func cleanup(res *http.Response, req *http.Request, apiTest *apitest.APITest) {
	if http.StatusOK == res.StatusCode || http.StatusCreated == res.StatusCode {
		configDB := _dbDriver.ConfigDB{
			DB_USERNAME: util.GetConfig("DB_USERNAME"),
			DB_PASSWORD: util.GetConfig("DB_PASSWORD"),
			DB_HOST:     util.GetConfig("DB_HOST"),
			DB_PORT:     util.GetConfig("DB_PORT"),
			DB_NAME:     util.GetConfig("DB_TEST_NAME"),
		}

		db := configDB.InitDB()

		_dbDriver.CleanSeeders(db)
	}
}

func getJWTToken(t *testing.T) string {
	configDB := _dbDriver.ConfigDB{
		DB_USERNAME: util.GetConfig("DB_USERNAME"),
		DB_PASSWORD: util.GetConfig("DB_PASSWORD"),
		DB_HOST:     util.GetConfig("DB_HOST"),
		DB_PORT:     util.GetConfig("DB_PORT"),
		DB_NAME:     util.GetConfig("DB_TEST_NAME"),
	}

	db := configDB.InitDB()

	user := _dbDriver.SeedUser(db)

	var userRequest *request.User = &request.User{
		Email:    user.Email,
		Password: user.Password,
	}

	var resp *http.Response = apitest.New().
		Handler(newApp()).
		Post("/api/v1/users/login").
		JSON(userRequest).
		Expect(t).
		Status(http.StatusOK).
		End().Response

	var response map[string]string = map[string]string{}

	json.NewDecoder(resp.Body).Decode(&response)

	var token string = response["token"]

	var JWT_TOKEN = "Bearer " + token

	return JWT_TOKEN
}

func getUser() users.User {
	configDB := _dbDriver.ConfigDB{
		DB_USERNAME: util.GetConfig("DB_USERNAME"),
		DB_PASSWORD: util.GetConfig("DB_PASSWORD"),
		DB_HOST:     util.GetConfig("DB_HOST"),
		DB_PORT:     util.GetConfig("DB_PORT"),
		DB_NAME:     util.GetConfig("DB_TEST_NAME"),
	}

	db := configDB.InitDB()

	user := _dbDriver.SeedUser(db)

	return user
}

func getNote() notes.Note {
	configDB := _dbDriver.ConfigDB{
		DB_USERNAME: util.GetConfig("DB_USERNAME"),
		DB_PASSWORD: util.GetConfig("DB_PASSWORD"),
		DB_HOST:     util.GetConfig("DB_HOST"),
		DB_PORT:     util.GetConfig("DB_PORT"),
		DB_NAME:     util.GetConfig("DB_TEST_NAME"),
	}

	db := configDB.InitDB()

	note := _dbDriver.SeedNote(db)

	return note
}

func getCategory() categories.Category {
	configDB := _dbDriver.ConfigDB{
		DB_USERNAME: util.GetConfig("DB_USERNAME"),
		DB_PASSWORD: util.GetConfig("DB_PASSWORD"),
		DB_HOST:     util.GetConfig("DB_HOST"),
		DB_PORT:     util.GetConfig("DB_PORT"),
		DB_NAME:     util.GetConfig("DB_TEST_NAME"),
	}

	db := configDB.InitDB()

	category := _dbDriver.SeedCategory(db)

	return category
}

func TestRegister_Success(t *testing.T) {
	var userRequest *request.User = &request.User{
		Email:    "test@mail.com",
		Password: "123123",
	}

	apitest.New().
		Observe(cleanup).
		Handler(newApp()).
		Post("/api/v1/users/register").
		JSON(userRequest).
		Expect(t).
		Status(http.StatusCreated).
		End()
}

func TestRegister_ValidationFailed(t *testing.T) {
	var userRequest *request.User = &request.User{
		Email:    "",
		Password: "",
	}

	apitest.New().
		Handler(newApp()).
		Post("/api/v1/users/register").
		JSON(userRequest).
		Expect(t).
		Status(http.StatusBadRequest).
		End()
}

func TestLogin_Success(t *testing.T) {
	user := getUser()

	var userRequest *request.User = &request.User{
		Email:    user.Email,
		Password: user.Password,
	}

	apitest.New().
		Handler(newApp()).
		Post("/api/v1/users/login").
		JSON(userRequest).
		Expect(t).
		Status(http.StatusOK).
		End()
}

func TestLogin_ValidationFailed(t *testing.T) {
	var userRequest *request.User = &request.User{
		Email:    "",
		Password: "",
	}

	apitest.New().
		Handler(newApp()).
		Post("/api/v1/users/login").
		JSON(userRequest).
		Expect(t).
		Status(http.StatusBadRequest).
		End()
}

func TestLogin_Failed(t *testing.T) {
	var userRequest *request.User = &request.User{
		Email:    "notfound@mail.com",
		Password: "123123",
	}

	apitest.New().
		Handler(newApp()).
		Post("/api/v1/users/login").
		JSON(userRequest).
		Expect(t).
		Status(http.StatusUnauthorized).
		End()
}

func TestGetNotes_Success(t *testing.T) {
	var token string = getJWTToken(t)

	apitest.New().
		Observe(cleanup).
		Handler(newApp()).
		Get("/api/v1/notes").
		Header("Authorization", token).
		Expect(t).
		Status(http.StatusOK).
		End()
}

func TestGetNote_Success(t *testing.T) {
	var note notes.Note = getNote()

	noteID := strconv.Itoa(int(note.ID))

	var token string = getJWTToken(t)

	apitest.New().
		Observe(cleanup).
		Handler(newApp()).
		Get("/api/v1/notes/"+noteID).
		Header("Authorization", token).
		Expect(t).
		Status(http.StatusOK).
		End()
}

func TestGetNote_NotFound(t *testing.T) {
	var token string = getJWTToken(t)

	apitest.New().
		Handler(newApp()).
		Get("/api/v1/notes/0").
		Header("Authorization", token).
		Expect(t).
		Status(http.StatusNotFound).
		End()
}

func TestCreateNote_Success(t *testing.T) {
	category := getCategory()

	var noteRequest *notes.Note = &notes.Note{
		Title:      "test",
		Content:    "test",
		CategoryID: category.ID,
	}

	var token string = getJWTToken(t)

	apitest.New().
		Observe(cleanup).
		Handler(newApp()).
		Post("/api/v1/notes").
		Header("Authorization", token).
		JSON(noteRequest).
		Expect(t).
		Status(http.StatusCreated).
		End()
}

func TestCreateNote_ValidationFailed(t *testing.T) {
	var noteRequest *notes.Note = &notes.Note{}

	var token string = getJWTToken(t)

	apitest.New().
		Handler(newApp()).
		Post("/api/v1/notes").
		Header("Authorization", token).
		JSON(noteRequest).
		Expect(t).
		Status(http.StatusBadRequest).
		End()
}

func TestUpdateNote_Success(t *testing.T) {
	var note notes.Note = getNote()

	category := getCategory()

	var noteRequest *notes.Note = &notes.Note{
		Title:      "test",
		Content:    "test",
		CategoryID: category.ID,
	}

	noteID := strconv.Itoa(int(note.ID))

	var token string = getJWTToken(t)

	apitest.New().
		Observe(cleanup).
		Handler(newApp()).
		Put("/api/v1/notes/"+noteID).
		Header("Authorization", token).
		JSON(noteRequest).
		Expect(t).
		Status(http.StatusOK).
		End()
}

func TestUpdateNote_ValidationFailed(t *testing.T) {
	var note notes.Note = getNote()

	var noteRequest *notes.Note = &notes.Note{}

	noteID := strconv.Itoa(int(note.ID))

	var token string = getJWTToken(t)

	apitest.New().
		Handler(newApp()).
		Put("/api/v1/notes/"+noteID).
		Header("Authorization", token).
		JSON(noteRequest).
		Expect(t).
		Status(http.StatusBadRequest).
		End()
}

func TestDeleteNote_Success(t *testing.T) {
	var note notes.Note = getNote()

	var token string = getJWTToken(t)

	noteID := strconv.Itoa(int(note.ID))

	apitest.New().
		Observe(cleanup).
		Handler(newApp()).
		Delete("/api/v1/notes/"+noteID).
		Header("Authorization", token).
		Expect(t).
		Status(http.StatusOK).
		End()
}

func TestDeleteNote_Failed(t *testing.T) {
	var token string = getJWTToken(t)

	apitest.New().
		Handler(newApp()).
		Observe(cleanup).
		Delete("/api/v1/notes/-1").
		Header("Authorization", token).
		Expect(t).
		Status(http.StatusInternalServerError).
		End()
}

func TestLogout_Success(t *testing.T) {
	var token string = getJWTToken(t)

	apitest.New().
		Handler(newApp()).
		Observe(cleanup).
		Post("/api/v1/users/logout").
		Header("Authorization", token).
		Expect(t).
		Status(http.StatusOK).
		End()
}

func TestLogout_Failed(t *testing.T) {
	apitest.New().
		Handler(newApp()).
		Observe(cleanup).
		Post("/api/v1/users/logout").
		Expect(t).
		Status(http.StatusBadRequest).
		End()
}
