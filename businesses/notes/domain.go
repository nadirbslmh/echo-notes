package notes

import (
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	ID           uint
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt
	Title        string
	Content      string
	CategoryName string
	CategoryID   uint
}

type Usecase interface {
	GetAll() []Domain
	GetByID(id string) Domain
	Create(noteDomain *Domain) Domain
	Update(id string, noteDomain *Domain) Domain
	Delete(id string) bool
	Restore(id string) Domain
	ForceDelete(id string) bool
}

type Repository interface {
	GetAll() []Domain
	GetByID(id string) Domain
	Create(noteDomain *Domain) Domain
	Update(id string, noteDomain *Domain) Domain
	Delete(id string) bool
	Restore(id string) Domain
	ForceDelete(id string) bool
}
