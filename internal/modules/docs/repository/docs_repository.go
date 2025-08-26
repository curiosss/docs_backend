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
func (r *DocsRepository) CreateDoc(doc *models.Doc) (*models.Doc, error) {
	if err := r.db.Create(doc).Error; err != nil {
		return nil, err
	}
	return doc, nil
}

func (r *DocsRepository) CreateDocUsers(docUsers []models.DocUser) error {
	return r.db.Create(docUsers).Error
}

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
