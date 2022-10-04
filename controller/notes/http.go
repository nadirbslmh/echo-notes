package notes

import (
	"echo-notes/businesses/notes"
	"echo-notes/controller/notes/request"
	"echo-notes/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

type NoteController struct {
	noteUseCase notes.Usecase
}

func NewNoteController(noteUC notes.Usecase) *NoteController {
	return &NoteController{
		noteUseCase: noteUC,
	}
}

func (ctrl *NoteController) GetAll(c echo.Context) error {
	notesData := ctrl.noteUseCase.GetAll()

	return c.JSON(http.StatusOK, model.Response[[]notes.Domain]{
		Status:  "success",
		Message: "all notes",
		Data:    notesData,
	})
}

func (ctrl *NoteController) GetByID(c echo.Context) error {
	var id string = c.Param("id")

	note := ctrl.noteUseCase.GetByID(id)

	if note.ID == 0 {
		return c.JSON(http.StatusNotFound, model.Response[string]{
			Status:  "failed",
			Message: "note not found",
		})
	}

	return c.JSON(http.StatusOK, model.Response[notes.Domain]{
		Status:  "success",
		Message: "note found",
		Data:    note,
	})
}

func (ctrl *NoteController) Create(c echo.Context) error {
	input := request.Note{}

	if err := c.Bind(&input); err != nil {
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

	note := ctrl.noteUseCase.Create(input.ToDomain())

	return c.JSON(http.StatusCreated, model.Response[notes.Domain]{
		Status:  "success",
		Message: "note created",
		Data:    note,
	})
}

func (ctrl *NoteController) Update(c echo.Context) error {
	input := request.Note{}

	if err := c.Bind(&input); err != nil {
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

	note := ctrl.noteUseCase.Update(noteId, input.ToDomain())

	if note.ID == 0 {
		return c.JSON(http.StatusNotFound, model.Response[string]{
			Status:  "failed",
			Message: "note not found",
		})
	}

	return c.JSON(http.StatusOK, model.Response[notes.Domain]{
		Status:  "success",
		Message: "note updated",
		Data:    note,
	})
}

func (ctrl *NoteController) Delete(c echo.Context) error {
	var noteId string = c.Param("id")

	isSuccess := ctrl.noteUseCase.Delete(noteId)

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

func (ctrl *NoteController) Restore(c echo.Context) error {
	var noteId string = c.Param("id")

	note := ctrl.noteUseCase.Restore(noteId)

	return c.JSON(http.StatusOK, model.Response[notes.Domain]{
		Status:  "success",
		Message: "data restored",
		Data:    note,
	})
}

func (ctrl *NoteController) ForceDelete(c echo.Context) error {
	var noteId string = c.Param("id")

	isSuccess := ctrl.noteUseCase.ForceDelete(noteId)

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
