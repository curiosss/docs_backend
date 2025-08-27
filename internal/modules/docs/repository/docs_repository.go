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
func (r *DocsRepository) GetDocsForUser(getDocsDto dto.GetDocsDto) (*dto.DocsResponseDto, error) {

	offset := (getDocsDto.Page - 1) * getDocsDto.Limit
	var docs []dto.DocResponse
	query := r.db.Table("docs").
		Select(`
		docs.*,
		doc_users.permission AS user_perm,
		users.username AS username,
		categories.name AS category_name
	`).Joins("LEFT JOIN doc_users ON doc_users.doc_id = docs.id AND doc_users.user_id = ?", getDocsDto.UserId).
		Joins("LEFT JOIN users ON users.id = docs.user_id").
		Joins(" LEFT JOIN categories ON categories.id = docs.category_id").
		Where("doc_users.user_id = ? OR docs.permission > 0", getDocsDto.UserId)

	if getDocsDto.CategoryID != nil {
		query = query.Where("docs.category_id = ?", *getDocsDto.CategoryID)
	}

	if err := query.Limit(getDocsDto.Limit).Offset(offset).Scan(&docs).Error; err != nil {
		return nil, err
	}

	res := &dto.DocsResponseDto{
		Docs:  docs,
		Total: 0,
	}
	return res, nil
}
