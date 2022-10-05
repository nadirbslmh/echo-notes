package categories_test

import (
	"echo-notes/businesses/categories"
	_categoryMock "echo-notes/businesses/categories/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	categoryRepository _categoryMock.Repository
	categoryService    categories.Usecase

	categoryDomain categories.Domain
)

func TestMain(m *testing.M) {
	categoryService = categories.NewCategoryUsecase(&categoryRepository)
	categoryDomain = categories.Domain{
		Name: "test",
	}
}

func TestGetAll(t *testing.T) {
	t.Run("Get All | Valid", func(t *testing.T) {
		categoryRepository.On("GetAll").Return([]categories.Domain{categoryDomain}).Once()

		result := categoryService.GetAll()

		assert.Equal(t, 0, len(result))
	})

	t.Run("Get All | InValid", func(t *testing.T) {
		categoryRepository.On("GetAll").Return([]categories.Domain{}).Once()

		result := categoryService.GetAll()

		assert.Equal(t, 0, len(result))
	})
}
