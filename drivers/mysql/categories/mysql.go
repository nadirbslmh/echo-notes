package categories

import (
	"echo-notes/businesses/categories"

	"gorm.io/gorm"
)

type categoryRepository struct {
	conn *gorm.DB
}

func NewMySQLRepository(conn *gorm.DB) categories.Repository {
	return &categoryRepository{
		conn: conn,
	}
}

func (cr *categoryRepository) GetAll() []categories.Domain {
	var rec []Category

	cr.conn.Find(&rec)

	categoryDomain := []categories.Domain{}

	for _, category := range rec {
		categoryDomain = append(categoryDomain, category.ToDomain())
	}

	return categoryDomain
}

func (cr *categoryRepository) GetByID(id string) categories.Domain {
	var category Category

	cr.conn.First(&category, "id = ?", id)

	return category.ToDomain()
}

func (cr *categoryRepository) Create(categoryDomain *categories.Domain) categories.Domain {
	rec := FromDomain(categoryDomain)

	result := cr.conn.Create(&rec)

	result.Last(&rec)

	return rec.ToDomain()
}

func (cr *categoryRepository) Update(id string, categoryDomain *categories.Domain) categories.Domain {
	var category categories.Domain = cr.GetByID(id)

	updatedCategory := FromDomain(&category)

	updatedCategory.Name = categoryDomain.Name

	cr.conn.Save(&updatedCategory)

	return updatedCategory.ToDomain()
}

func (cr *categoryRepository) Delete(id string) bool {
	var category categories.Domain = cr.GetByID(id)

	deletedCategory := FromDomain(&category)

	result := cr.conn.Unscoped().Delete(deletedCategory)

	if result.RowsAffected == 0 {
		return false
	}

	return true
}
