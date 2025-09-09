package repository

import (
	"docs-notify/internal/models"
	"docs-notify/internal/modules/notifications/dto"
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

func (r *NotifsRepository) GetAll(userId uint, page int) ([]models.Notification, error) {

	offset := (page - 1) * 30

	var items []models.Notification

	err := r.db.
		Where("user_id = ?", userId).
		Limit(30).
		Offset(offset).
		Find(&items).Error

	if err != nil {
		return nil, err
	}
	return items, nil
}

func (r *NotifsRepository) GetAdminNotifs(docId uint) ([]dto.NotificationAdminResponseDto, error) {

	var items []dto.NotificationAdminResponseDto
	query := r.db.Table("notifications").
		Select(`
		notifications.*,
		users.username AS username,
		users.role as user_role
	`).Joins("LEFT JOIN users ON users.id = notifications.user_id").
		Where("notifications.doc_id = ?", docId)

	if err := query.Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

func (r *NotifsRepository) MarkAsSeen(userId uint, docId uint) error {
	err := r.db.Model(&models.Notification{}).Where("doc_id = ? AND user_id = ?", docId, userId).Update("is_seen", true).Error
	if err != nil {
		return err
	}
	return nil
}

// func (r *NotifsRepository) Update(category *models.Notification) (*models.Category, error) {
// 	// Ensure the category exists before updating
// 	var existing models.Category
// 	if err := r.db.First(&existing, category.ID).Error; err != nil {
// 		return nil, errors.New("kategoriýa tapylmady")
// 	}

// 	// Perform update
// 	if err := r.db.Model(&existing).Updates(category).Error; err != nil {
// 		return nil, err
// 	}

// 	return &existing, nil
// }

// func (r *NotifsRepository) Delete(id uint) error {
// 	result := r.db.Delete(&models.Category{}, id)

// 	if result.RowsAffected > 0 {
// 		return nil
// 	} else {
// 		if result.Error == nil {
// 			return errors.New("Kategoriýa tapylmady")
// 		}
// 		return result.Error
// 	}
// }
