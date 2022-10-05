package notes

import (
	"echo-notes/businesses/notes"
	"echo-notes/controller"
	"echo-notes/controller/notes/request"
	"echo-notes/controller/notes/response"
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

	notes := []response.Note{}

	for _, note := range notesData {
		notes = append(notes, response.FromDomain(note))
	}

	return controller.NewResponse(c, http.StatusOK, "success", "all notes", notes)
}

func (ctrl *NoteController) GetByID(c echo.Context) error {
	var id string = c.Param("id")

	note := ctrl.noteUseCase.GetByID(id)

	if note.ID == 0 {
		return controller.NewResponse(c, http.StatusNotFound, "failed", "note not found", "")
	}

	return controller.NewResponse(c, http.StatusOK, "success", "note found", response.FromDomain(note))
}

func (ctrl *NoteController) Create(c echo.Context) error {
	input := request.Note{}

	if err := c.Bind(&input); err != nil {
		return controller.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}

	err := input.Validate()

	if err != nil {
		return controller.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}

	note := ctrl.noteUseCase.Create(input.ToDomain())

	return controller.NewResponse(c, http.StatusCreated, "success", "note created", response.FromDomain(note))
}

func (ctrl *NoteController) Update(c echo.Context) error {
	input := request.Note{}

	if err := c.Bind(&input); err != nil {
		return controller.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}

	var noteId string = c.Param("id")

	err := input.Validate()

	if err != nil {
		return controller.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}

	note := ctrl.noteUseCase.Update(noteId, input.ToDomain())

	if note.ID == 0 {
		return controller.NewResponse(c, http.StatusNotFound, "failed", "note not found", "")
	}

	return controller.NewResponse(c, http.StatusOK, "success", "note updated", response.FromDomain(note))
}

func (ctrl *NoteController) Delete(c echo.Context) error {
	var noteId string = c.Param("id")

	isSuccess := ctrl.noteUseCase.Delete(noteId)

	if !isSuccess {
		return controller.NewResponse(c, http.StatusNotFound, "failed", "note not found", "")
	}

	return controller.NewResponse(c, http.StatusOK, "success", "note deleted", "")
}

func (ctrl *NoteController) Restore(c echo.Context) error {
	var noteId string = c.Param("id")

	note := ctrl.noteUseCase.Restore(noteId)

	if note.ID == 0 {
		return controller.NewResponse(c, http.StatusNotFound, "failed", "note not found", "")
	}

	return controller.NewResponse(c, http.StatusOK, "success", "note restored", response.FromDomain(note))
}

func (ctrl *NoteController) ForceDelete(c echo.Context) error {
	var noteId string = c.Param("id")

	isSuccess := ctrl.noteUseCase.ForceDelete(noteId)

	if !isSuccess {
		return controller.NewResponse(c, http.StatusNotFound, "failed", "note not found", "")
	}

	return controller.NewResponse(c, http.StatusOK, "success", "note deleted permanently", "")
}
