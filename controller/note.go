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

	return c.JSON(http.StatusOK, model.Response[[]model.Note]{
		Status:  "success",
		Message: "all notes",
		Data:    notes,
	})
}

func GetByID(c echo.Context) error {
	var id string = c.Param("id")

	note := noteService.GetByID(id)

	if note.ID == 0 {
		return c.JSON(http.StatusNotFound, model.Response[string]{
			Status:  "failed",
			Message: "note not found",
		})
	}

	return c.JSON(http.StatusOK, model.Response[model.Note]{
		Status:  "success",
		Message: "note found",
		Data:    note,
	})
}

func Create(c echo.Context) error {
	var input *model.NoteInput = new(model.NoteInput)

	if err := c.Bind(input); err != nil {
		return c.JSON(http.StatusBadRequest, model.Response[string]{
			Status:  "failed",
			Message: "validation failed",
		})
	}

	err := input.Validate()

	if err != nil {
		return c.JSON(http.StatusBadRequest, model.Response[string]{
			Status:  "failed",
			Message: "validation failed",
		})
	}

	note := noteService.Create(*input)

	return c.JSON(http.StatusCreated, model.Response[model.Note]{
		Status:  "success",
		Message: "note created",
		Data:    note,
	})
}

func Update(c echo.Context) error {
	var input *model.NoteInput = new(model.NoteInput)

	if err := c.Bind(input); err != nil {
		return c.JSON(http.StatusBadRequest, model.Response[string]{
			Status:  "failed",
			Message: "validation failed",
		})
	}

	var noteId string = c.Param("id")

	err := input.Validate()

	if err != nil {
		return c.JSON(http.StatusBadRequest, model.Response[string]{
			Status:  "failed",
			Message: "validation failed",
		})
	}

	note := noteService.Update(noteId, *input)

	if note.ID == 0 {
		return c.JSON(http.StatusNotFound, model.Response[string]{
			Status:  "failed",
			Message: "note not found",
		})
	}

	return c.JSON(http.StatusOK, model.Response[model.Note]{
		Status:  "success",
		Message: "note updated",
		Data:    note,
	})
}

func Delete(c echo.Context) error {
	var noteId string = c.Param("id")

	isSuccess := noteService.Delete(noteId)

	if !isSuccess {
		return c.JSON(http.StatusInternalServerError, model.Response[string]{
			Status:  "failed",
			Message: "data deletion failed",
		})
	}

	return c.JSON(http.StatusOK, model.Response[string]{
		Status:  "success",
		Message: "note deleted",
	})
}

func Restore(c echo.Context) error {
	var noteId string = c.Param("id")

	note := noteService.Restore(noteId)

	return c.JSON(http.StatusOK, model.Response[model.Note]{
		Status:  "success",
		Message: "data restored",
		Data:    note,
	})
}

func ForceDelete(c echo.Context) error {
	var noteId string = c.Param("id")

	isSuccess := noteService.ForceDelete(noteId)

	if !isSuccess {
		return c.JSON(http.StatusInternalServerError, model.Response[string]{
			Status:  "failed",
			Message: "data force deletion failed",
		})
	}

	return c.JSON(http.StatusOK, model.Response[string]{
		Status:  "success",
		Message: "note force deleted",
	})
}
