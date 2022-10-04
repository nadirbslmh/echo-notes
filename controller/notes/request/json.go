package request

import (
	"echo-notes/businesses/notes"

	"github.com/go-playground/validator/v10"
)

type Note struct {
	Title      string `json:"title" validate:"required"`
	Content    string `json:"content" validate:"required"`
	CategoryID uint   `json:"category_id" validate:"required"`
}

func (req *Note) ToDomain() *notes.Domain {
	return &notes.Domain{
		Title:      req.Title,
		Content:    req.Content,
		CategoryID: req.CategoryID,
	}
}

func (req *Note) Validate() error {
	validate := validator.New()

	err := validate.Struct(req)

	return err
}
