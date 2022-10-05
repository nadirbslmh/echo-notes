package notes_test

import (
	"echo-notes/businesses/categories"
	"echo-notes/businesses/notes"
	_noteMock "echo-notes/businesses/notes/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	noteRepository _noteMock.Repository
	noteService    notes.Usecase

	noteDomain notes.Domain
)

func TestMain(m *testing.M) {
	noteService = notes.NewNoteUsecase(&noteRepository)

	categoryDomain := categories.Domain{
		Name: "test category",
	}

	noteDomain = notes.Domain{
		Title:      "title",
		Content:    "my content",
		CategoryID: categoryDomain.ID,
	}

	m.Run()
}

func TestGetAll(t *testing.T) {
	t.Run("Get All | Valid", func(t *testing.T) {
		noteRepository.On("GetAll").Return([]notes.Domain{noteDomain}).Once()

		result := noteService.GetAll()

		assert.Equal(t, 1, len(result))
	})

	t.Run("Get All | InValid", func(t *testing.T) {
		noteRepository.On("GetAll").Return([]notes.Domain{}).Once()

		result := noteService.GetAll()

		assert.Equal(t, 0, len(result))
	})
}

func TestGetByID(t *testing.T) {
	t.Run("Get By ID | Valid", func(t *testing.T) {
		noteRepository.On("GetByID", "1").Return(noteDomain).Once()

		result := noteService.GetByID("1")

		assert.NotNil(t, result)
	})

	t.Run("Get By ID | InValid", func(t *testing.T) {
		noteRepository.On("GetByID", "-1").Return(notes.Domain{}).Once()

		result := noteService.GetByID("-1")

		assert.NotNil(t, result)
	})
}

func TestCreate(t *testing.T) {
	t.Run("Create | Valid", func(t *testing.T) {
		noteRepository.On("Create", &noteDomain).Return(noteDomain).Once()

		result := noteService.Create(&noteDomain)

		assert.NotNil(t, result)
	})

	t.Run("Create | InValid", func(t *testing.T) {
		noteRepository.On("Create", &notes.Domain{}).Return(notes.Domain{}).Once()

		result := noteService.Create(&notes.Domain{})

		assert.NotNil(t, result)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Update | Valid", func(t *testing.T) {
		noteRepository.On("Update", "1", &noteDomain).Return(noteDomain).Once()

		result := noteService.Update("1", &noteDomain)

		assert.NotNil(t, result)
	})

	t.Run("Update | InValid", func(t *testing.T) {
		noteRepository.On("Update", "1", &notes.Domain{}).Return(notes.Domain{}).Once()

		result := noteService.Update("1", &notes.Domain{})

		assert.NotNil(t, result)
	})
}

func TestDelete(t *testing.T) {
	t.Run("Delete | Valid", func(t *testing.T) {
		noteRepository.On("Delete", "1").Return(true).Once()

		result := noteService.Delete("1")

		assert.True(t, result)
	})

	t.Run("Delete | InValid", func(t *testing.T) {
		noteRepository.On("Delete", "-1").Return(false).Once()

		result := noteService.Delete("-1")

		assert.False(t, result)
	})
}

func TestRestore(t *testing.T) {
	t.Run("Restore | Valid", func(t *testing.T) {
		noteRepository.On("Restore", "1").Return(noteDomain).Once()

		result := noteService.Restore("1")

		assert.NotNil(t, result)
	})

	t.Run("Restore | InValid", func(t *testing.T) {
		noteRepository.On("Restore", "-1").Return(notes.Domain{}).Once()

		result := noteService.Restore("-1")

		assert.NotNil(t, result)
	})
}

func TestForceDelete(t *testing.T) {
	t.Run("ForceDelete | Valid", func(t *testing.T) {
		noteRepository.On("ForceDelete", "1").Return(true).Once()

		result := noteService.ForceDelete("1")

		assert.True(t, result)
	})

	t.Run("ForceDelete | InValid", func(t *testing.T) {
		noteRepository.On("ForceDelete", "-1").Return(false).Once()

		result := noteService.ForceDelete("-1")

		assert.False(t, result)
	})
}
