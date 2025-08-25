package repository

import (
	"docs-notify/internal/models"

	"gorm.io/gorm"
)

type DocsRepository struct {
	db *gorm.DB
}

func NewDocsRepository(db *gorm.DB) *DocsRepository {
	return &DocsRepository{db: db}
}
func (r *DocsRepository) CreateDoc(doc *models.Doc) error {
	return r.db.Create(doc).Error
}

// func (r *DocsRepository) CreateFiles(files []models.File) error {
// 	return r.db.Create(&files).Error
// }
