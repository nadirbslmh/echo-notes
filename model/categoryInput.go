package model

import "github.com/go-playground/validator/v10"

type CategoryInput struct {
	Name string `json:"name" validate:"required"`
}

func (categoryInput *CategoryInput) Validate() error {
	validate := validator.New()

	err := validate.Struct(categoryInput)

	return err
}
