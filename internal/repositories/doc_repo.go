package repositories

import (
	"time"

	"docs-notify/internal/models"

	"gorm.io/gorm"
)

type DocRepository interface {
	Create(doc *models.Doc) error
	FindByID(id uint) (*models.Doc, error)
	FindForNotification() ([]models.Doc, error)
	UpdateNotifSent(id uint) error
	AddFile(file *models.File) error
}

type docRepository struct {
	db *gorm.DB
}

func NewDocRepository(db *gorm.DB) DocRepository {
	return &docRepository{db}
}

func (r *docRepository) Create(doc *models.Doc) error {
	return r.db.Create(doc).Error
}

func (r *docRepository) FindByID(id uint) (*models.Doc, error) {
	var doc models.Doc
	if err := r.db.Preload("Author").Preload("Category").First(&doc, id).Error; err != nil {
		return nil, err
	}
	return &doc, nil
}

func (r *docRepository) FindForNotification() ([]models.Doc, error) {
	var docs []models.Doc
	err := r.db.Where("notify_date <= ? AND notif_sent = ?", time.Now(), false).Find(&docs).Error
	return docs, err
}

func (r *docRepository) UpdateNotifSent(id uint) error {
	return r.db.Model(&models.Doc{}).Where("id = ?", id).Update("notif_sent", true).Error
}

func (r *docRepository) AddFile(file *models.File) error {
	return r.db.Create(file).Error
}
