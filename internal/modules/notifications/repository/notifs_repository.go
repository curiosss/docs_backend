package repository

import (
	"docs-notify/internal/models"
	"errors"

	"gorm.io/gorm"
)

type NotifsRepository struct {
	db *gorm.DB
}

func NewNotifsRepository(db *gorm.DB) *NotifsRepository {
	return &NotifsRepository{db: db}
}

func (r *NotifsRepository) Add(notification *models.Notification) (*models.Notification, error) {

	//Check for existance
	var existing models.Notification
	err := r.db.Where("doc_id = ? AND user_id=?", notification.DocID, notification.UserID).First(&existing).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if existing.ID > 0 {
		//Update
		notification.ID = existing.ID
		if err = r.db.Save(notification).Error; err != nil {
			return nil, err
		}
		return notification, nil
	} else {
		// Save to database
		if err = r.db.Create(notification).Error; err != nil {
			return nil, err
		}
		return notification, nil
	}
}

func (r *NotifsRepository) Update(category *models.Notification) (*models.Category, error) {
	// Ensure the category exists before updating
	var existing models.Category
	if err := r.db.First(&existing, category.ID).Error; err != nil {
		return nil, errors.New("kategoriýa tapylmady")
	}

	// Perform update
	if err := r.db.Model(&existing).Updates(category).Error; err != nil {
		return nil, err
	}

	return &existing, nil
}

func (r *NotifsRepository) Delete(id uint) error {
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

func (r *NotifsRepository) GetAll() ([]models.Category, error) {
	var categories []models.Category

	// Fetch only necessary fields and order by role
	if err := r.db.Model(&models.Category{}).
		Scan(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}
