package repositories

import (
	"indocattes/models"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	FindCategorys() ([]models.Category, error)
	GetCategory(ID int) (models.Category, error)
	CreateCategory(category models.Category) (models.Category, error)
	UpdateCategory(category models.Category) (models.Category, error)
	DeleteCategory(category models.Category) (models.Category, error)
	// GetCategoryfilm(ID int) (models.Category, error)
}

func RepositoryCategory(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindCategorys() ([]models.Category, error) {
	var categorys []models.Category
	err := r.db.Find(&categorys).Error

	return categorys, err
}

func (r *repository) GetCategory(ID int) (models.Category, error) {
	var category models.Category
	err := r.db.First(&category, ID).Error

	return category, err
}

func (r *repository) CreateCategory(category models.Category) (models.Category, error) {
	err := r.db.Create(&category).Error

	return category, err
}

func (r *repository) UpdateCategory(category models.Category) (models.Category, error) {
	err := r.db.Model(&category).Updates(&category).Error

	return category, err
}

func (r *repository) DeleteCategory(category models.Category) (models.Category, error) {
	err := r.db.Delete(&category).Error

	return category, err
}
