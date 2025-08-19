package repository

import (
	"docs-notify/internal/models"
	"docs-notify/internal/modules/categories/dto"
	"errors"

	"gorm.io/gorm"
)

type CategoriesRepository struct {
	db *gorm.DB
}

func NewCategRepository(db *gorm.DB) *CategoriesRepository {
	return &CategoriesRepository{db: db}
}

func (r *CategoriesRepository) Create(categoryDto *dto.CategoryCreateDto) (*models.Category, error) {
	// Map DTO to model
	category := &models.Category{
		Name:     categoryDto.Name,
		ParentID: categoryDto.ParentID,
	}

	// // Check if a category with the same name already exists
	// var existing models.Category
	// if err := r.db.Where("name = ?", categoryDto.Name).First(&existing).Error; err == nil {
	// 	return nil, errors.New("Bu Kategoriýa ýa-da Subkategoriýa ady öňden bar")
	// } else if !errors.Is(err, gorm.ErrRecordNotFound) {
	// 	return nil, err
	// }

	// Save to database
	if err := r.db.Create(category).Error; err != nil {
		return nil, err
	}

	return category, nil
}

func (r *CategoriesRepository) Update(category *models.Category) (*models.Category, error) {
	// Ensure the category exists before updating
	var existing models.Category
	if err := r.db.First(&existing, category.ID).Error; err != nil {
		return nil, errors.New("Kategoriýa tapylmady")
	}

	// Perform update
	if err := r.db.Model(&existing).Updates(category).Error; err != nil {
		return nil, err
	}

	return &existing, nil
}

func (r *CategoriesRepository) Delete(id uint) error {
	result := r.db.Delete(&models.Category{}, id)

	if result.RowsAffected > 0 {
		return nil
	} else {
		if result.Error == nil {
			return errors.New("Kategoriýa tapylmady")
		}
		return result.Error
	}
}

func (r *CategoriesRepository) GetAll() ([]models.Category, error) {
	var categories []models.Category

	// Fetch only necessary fields and order by role
	if err := r.db.Model(&models.Category{}).
		Scan(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}
