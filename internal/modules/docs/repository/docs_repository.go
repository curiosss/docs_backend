package repository

import (
	"docs-notify/internal/models"
	"docs-notify/internal/modules/docs/dto"

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

// GetDocsForUser returns docs for a user (explicit + global permission), with pagination
func (r *DocsRepository) GetDocsForUser(userId uint, page, pageSize int) (*dto.DocsResponseDto, error) {
	var docs []models.Doc
	var total int64

	offset := (page - 1) * pageSize

	// Build query
	query := r.db.Table("docs").
		Select("DISTINCT docs.*").
		Joins("LEFT JOIN doc_users ON doc_users.doc_id = docs.id").
		Where("doc_users.user_id = ? OR docs.permission > 0", userId)

	// Count total
	if err := query.Count(&total).Error; err != nil {
		return nil, err
	}

	// Apply pagination
	if err := query.Limit(pageSize).Offset(offset).Find(&docs).Error; err != nil {
		return nil, err
	}

	res := &dto.DocsResponseDto{
		Data:  docs,
		Total: total,
	}
	return res, nil
}
