package notes

import (
	noteUsecase "echo-notes/businesses/notes"
	"echo-notes/drivers/mysql/categories"
	"time"

	"gorm.io/gorm"
)

type Note struct {
	ID         uint                `json:"id" gorm:"primaryKey"`
	CreatedAt  time.Time           `json:"created_at"`
	UpdatedAt  time.Time           `json:"updated_at"`
	DeletedAt  gorm.DeletedAt      `json:"deleted_at"`
	Title      string              `json:"title"`
	Content    string              `json:"content"`
	Category   categories.Category `json:"category"`
	CategoryID uint                `json:"category_id"`
}

func FromDomain(domain *noteUsecase.Domain) *Note {
	return &Note{
		ID:         domain.ID,
		Title:      domain.Title,
		Content:    domain.Content,
		CategoryID: domain.CategoryID,
		CreatedAt:  domain.CreatedAt,
		UpdatedAt:  domain.UpdatedAt,
		DeletedAt:  domain.DeletedAt,
	}
}

func (rec *Note) ToDomain() noteUsecase.Domain {
	return noteUsecase.Domain{
		ID:           rec.ID,
		Title:        rec.Title,
		Content:      rec.Content,
		CategoryName: rec.Category.Name,
		CategoryID:   rec.Category.ID,
		CreatedAt:    rec.CreatedAt,
		UpdatedAt:    rec.UpdatedAt,
		DeletedAt:    rec.DeletedAt,
	}
}
