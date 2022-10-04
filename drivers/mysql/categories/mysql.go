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

	result := cr.conn.Create(&categoryDomain)

	var createdCategory categories.Domain

	result.Last(&createdCategory)

	return createdCategory
}

func (cr *categoryRepository) Update(id string, categoryDomain *categories.Domain) categories.Domain {
	var category categories.Domain = cr.GetByID(id)

	category.Name = categoryDomain.Name

	cr.conn.Save(&category)

	return category
}

func (cr *categoryRepository) Delete(id string) bool {
	var category categories.Domain = cr.GetByID(id)

	result := cr.conn.Unscoped().Delete(&category)

	if result.RowsAffected == 0 {
		return false
	}

	return true
}
