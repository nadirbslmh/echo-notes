package controller

import (
	"bytes"
	"echo-notes/database"
	"echo-notes/model"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func InitEcho() *echo.Echo {
	database.InitTestDB()
	e := echo.New()

	return e
}

func TestGetAllNotes_Success(t *testing.T) {
	var testCases = []struct {
		name                   string
		path                   string
		expectedStatus         int
		expectedBodyStartsWith string
	}{{
		name:                   "success",
		path:                   "/api/v1/notes",
		expectedStatus:         http.StatusOK,
		expectedBodyStartsWith: "{\"status\":",
	},
	}

	e := InitEcho()

	req := httptest.NewRequest(http.MethodGet, "/api/v1/notes", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)

		if assert.NoError(t, GetAll(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			body := rec.Body.String()

			assert.True(t, strings.HasPrefix(body, testCase.expectedBodyStartsWith))
		}
	}
}

func TestCreateNote_Success(t *testing.T) {
	var testCases = []struct {
		name                   string
		path                   string
		expectedStatus         int
		expectedBodyStartsWith string
	}{{
		name:                   "success",
		path:                   "/api/v1/notes",
		expectedStatus:         http.StatusCreated,
		expectedBodyStartsWith: "{\"status\":",
	},
	}

	e := InitEcho()

	category := database.SeedCategory()

	noteInput := model.NoteInput{
		Title:      "test",
		Content:    "test",
		CategoryID: category.ID,
	}

	jsonBody, _ := json.Marshal(&noteInput)
	bodyReader := bytes.NewReader(jsonBody)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/notes", bodyReader)
	rec := httptest.NewRecorder()

	req.Header.Add("Content-Type", "application/json")

	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)

		if assert.NoError(t, Create(c)) {
			assert.Equal(t, http.StatusCreated, rec.Code)
			body := rec.Body.String()

			assert.True(t, strings.HasPrefix(body, testCase.expectedBodyStartsWith))
		}
	}

}

func TestGetNoteByID_Success(t *testing.T) {
	var testCases = []struct {
		name                   string
		path                   string
		expectedStatus         int
		expectedBodyStartsWith string
	}{{
		name:                   "success",
		path:                   "/api/v1/notes",
		expectedStatus:         http.StatusOK,
		expectedBodyStartsWith: "{\"status\":",
	},
	}

	e := InitEcho()

	note := database.SeedNote()
	noteID := strconv.Itoa(int(note.ID))

	req := httptest.NewRequest(http.MethodGet, "/api/v1/notes", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)
		c.SetParamNames("id")
		c.SetParamValues(noteID)

		if assert.NoError(t, GetByID(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			body := rec.Body.String()

			assert.True(t, strings.HasPrefix(body, testCase.expectedBodyStartsWith))
		}
	}

}

func TestUpdateNote_Success(t *testing.T) {
	var testCases = []struct {
		name                   string
		path                   string
		expectedStatus         int
		expectedBodyStartsWith string
	}{{
		name:                   "success",
		path:                   "/api/v1/notes",
		expectedStatus:         http.StatusOK,
		expectedBodyStartsWith: "{\"status\":",
	},
	}

	e := InitEcho()

	note := database.SeedNote()

	noteInput := model.NoteInput{
		Title:      "test",
		Content:    "test",
		CategoryID: note.CategoryID,
	}

	jsonBody, _ := json.Marshal(&noteInput)
	bodyReader := bytes.NewReader(jsonBody)

	noteID := strconv.Itoa(int(note.ID))

	req := httptest.NewRequest(http.MethodPut, "/api/v1/notes", bodyReader)
	rec := httptest.NewRecorder()

	req.Header.Add("Content-Type", "application/json")

	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)
		c.SetParamNames("id")
		c.SetParamValues(noteID)

		if assert.NoError(t, Update(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			body := rec.Body.String()

			assert.True(t, strings.HasPrefix(body, testCase.expectedBodyStartsWith))
		}
	}

}

func TestDeleteNote_Success(t *testing.T) {
	var testCases = []struct {
		name                   string
		path                   string
		expectedStatus         int
		expectedBodyStartsWith string
	}{{
		name:                   "success",
		path:                   "/api/v1/notes",
		expectedStatus:         http.StatusOK,
		expectedBodyStartsWith: "{\"status\":",
	},
	}

	e := InitEcho()

	note := database.SeedNote()
	noteID := strconv.Itoa(int(note.ID))

	req := httptest.NewRequest(http.MethodDelete, "/api/v1/notes", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)
		c.SetParamNames("id")
		c.SetParamValues(noteID)

		if assert.NoError(t, Delete(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			body := rec.Body.String()

			assert.True(t, strings.HasPrefix(body, testCase.expectedBodyStartsWith))
		}
	}

}
