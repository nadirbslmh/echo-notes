package main

import (
	"echo-notes/database"
	"echo-notes/model"
	"echo-notes/route"
	"encoding/json"
	"net/http"
	"strconv"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/steinfletcher/apitest"
)

func newApp() *echo.Echo {
	database.InitTestDB()

	app := echo.New()

	route.SetupRoute(app)

	return app
}

func cleanup(res *http.Response, req *http.Request, apiTest *apitest.APITest) {
	if http.StatusOK == res.StatusCode || http.StatusCreated == res.StatusCode {
		database.CleanSeeders()
	}
}

func getJWTToken(t *testing.T) string {
	user := database.SeedUser()

	var userRequest *model.UserInput = &model.UserInput{
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

func TestRegister_Success(t *testing.T) {
	var userRequest *model.UserInput = &model.UserInput{
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
	var userRequest *model.UserInput = &model.UserInput{
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
	user := database.SeedUser()

	var userRequest *model.UserInput = &model.UserInput{
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
	var userRequest *model.UserInput = &model.UserInput{
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
	var userRequest *model.UserInput = &model.UserInput{
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
	var note model.Note = database.SeedNote()

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
	category := database.SeedCategory()

	var noteRequest *model.NoteInput = &model.NoteInput{
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
	var noteRequest *model.NoteInput = &model.NoteInput{}

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
	var note model.Note = database.SeedNote()

	category := database.SeedCategory()

	var noteRequest *model.NoteInput = &model.NoteInput{
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
	var note model.Note = database.SeedNote()

	var noteRequest *model.NoteInput = &model.NoteInput{}

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
	var note model.Note = database.SeedNote()

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
