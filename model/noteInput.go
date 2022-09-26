package model

import "github.com/go-playground/validator/v10"

type NoteInput struct {
	Title      string `json:"title" validate:"required"`
	Content    string `json:"content" validate:"required"`
	CategoryID uint   `json:"category_id" validate:"required"`
}

func (input *NoteInput) Validate() error {
	validate := validator.New()

	err := validate.Struct(input)

	return err
}
