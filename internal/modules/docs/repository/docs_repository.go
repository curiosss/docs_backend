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

//	func (r *DocsRepository) CreateFiles(files []models.File) error {
//		return r.db.Create(&files).Error
//	}
func (r *DocsRepository) GetDocs(userId uint) ([]models.Doc, error) {
	var docs []models.Doc

	err := r.db.Table("docs").
		Select("docs.*").
		Joins("JOIN doc_users ON doc_users.doc_id = docs.id").
		Where("doc_users.user_id = ?", userId).
		Find(&docs).Error

	if err != nil {
		return nil, err
	}
	return docs, nil
}
