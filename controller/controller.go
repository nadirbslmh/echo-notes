package controller

import (
	"echo-notes/model"
	"echo-notes/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

var noteService service.NoteService = service.New()

func GetAll(c echo.Context) error {
	var notes []model.Note = noteService.GetAll()

	return c.JSON(http.StatusOK, notes)
}

func GetByID(c echo.Context) error {
	var id string = c.Param("id")

	note := noteService.GetByID(id)

	if note.ID == 0 {
		return c.JSON(http.StatusNotFound, map[string]string{
			"message": "note not found",
		})
	}

	return c.JSON(http.StatusOK, note)
}

func Create(c echo.Context) error {
	var input *model.NoteInput = new(model.NoteInput)

	if err := c.Bind(input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "invalid request",
		})
	}

	note := noteService.Create(*input)

	return c.JSON(http.StatusCreated, note)
}

func Update(c echo.Context) error {
	var input *model.NoteInput = new(model.NoteInput)

	if err := c.Bind(input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "invalid request",
		})
	}

	var noteId string = c.Param("id")

	note := noteService.Update(noteId, *input)

	return c.JSON(http.StatusOK, note)
}

func Delete(c echo.Context) error {
	var noteId string = c.Param("id")

	isSuccess := noteService.Delete(noteId)

	if !isSuccess {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "failed to delete a data",
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "data deleted",
	})
}
