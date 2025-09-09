package repository

import (
	"docs-notify/internal/models"
	"docs-notify/internal/modules/docs/dto"
	"errors"
	"fmt"
	"log"
	"time"

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

	// Apply filters
	query = getFiltersQuery(getDocsDto, query)

	if err := query.Limit(getDocsDto.Limit).Offset(offset).Scan(&docs).Error; err != nil {
		return nil, err
	}

	res := &dto.DocsResponseDto{
		Docs:  docs,
		Total: 0,
	}
	return res, nil
}

func getFiltersQuery(getDocsDto dto.GetDocsDto, query *gorm.DB) *gorm.DB {
	if getDocsDto.Categories != nil && len(getDocsDto.Categories) > 0 {
		query = query.Where("docs.category_id IN ?", getDocsDto.Categories) // pass slice directly
	}
	if getDocsDto.CategoryID != nil {
		query = query.Where("docs.category_id = ?", *getDocsDto.CategoryID)
	}
	if getDocsDto.SubCategoryId != nil {
		query = query.Where("docs.sub_category_id = ?", *getDocsDto.SubCategoryId)
	}
	if getDocsDto.SearchText != nil && *getDocsDto.SearchText != "" {
		search := fmt.Sprintf("%%%s%%", *getDocsDto.SearchText)
		query = query.Where("docs.doc_name ILIKE ?", search)
	}
	if getDocsDto.Status != nil {
		query = query.Where("docs.status = ?", *getDocsDto.Status)
	}
	if getDocsDto.CreatedUserId != nil {
		query = query.Where("docs.user_id = ?", *getDocsDto.CreatedUserId)
	}

	// Update fields
	layout := "2006-01-02" // Go's reference date for parsing YYYY-MM-DD
	if getDocsDto.CreatedFrom != nil {
		createdFrom, err := time.Parse(layout, *getDocsDto.CreatedFrom)
		if err == nil {
			query = query.Where("docs.created_at >= ?", createdFrom)
		}

	}
	if getDocsDto.CreatedTo != nil {
		createdTo, err := time.Parse(layout, *getDocsDto.CreatedTo)
		if err == nil {
			query = query.Where("docs.created_at <= ?", createdTo)
		}

	}
	if getDocsDto.PreparedFrom != nil {
		preparedFrom, err := time.Parse(layout, *getDocsDto.PreparedFrom)
		if err == nil {
			query = query.Where("docs.prepared_date >= ?", preparedFrom)
		} else {
			log.Println("Error parsing prepared_from:", err)
		}

	}
	if getDocsDto.PreparedTo != nil {
		preparedTo, err := time.Parse(layout, *getDocsDto.PreparedTo)
		if err == nil {
			query = query.Where("docs.prepared_date <= ?", preparedTo)
		}

	}
	return query
}

func (r *DocsRepository) GetDocById(docId uint) (*dto.DocResponse, error) {
	var doc dto.DocResponse
	err := r.db.Table("docs").
		Select(`
		docs.*,
		users.username AS username
	`).Joins("LEFT JOIN users ON users.id = docs.user_id").
		Where("docs.id = ?", docId).Scan(&doc).Error

	if err != nil {
		return nil, err
	}

	return &doc, nil
}

func (r *DocsRepository) Delete(id uint) error {
	result := r.db.Delete(&models.Doc{}, id)
	fmt.Println(result.Error)
	if result.RowsAffected > 0 {
		return nil
	} else {
		if result.Error == nil {
			return errors.New("Dokument tapylmady")
		}
		return result.Error
	}
}

func (r *DocsRepository) DeleteDocUsersByDocID(docID uint) error {
	return r.db.Where("doc_id = ?", docID).Delete(&models.DocUser{}).Error
}

func (r *DocsRepository) DeleteDocNotifications(docID uint) error {
	return r.db.Where("doc_id=?", docID).Delete(&models.Notification{}).Error
}

func (r *DocsRepository) GetByID(id uint) (*models.Doc, error) {
	var doc models.Doc
	if err := r.db.First(&doc, id).Error; err != nil {
		return nil, err
	}
	return &doc, nil
}

func (r *DocsRepository) GetDocPermissions(docId uint) ([]models.DocUser, error) {
	var permissions []models.DocUser
	err := r.db.Model(&models.DocUser{}).Where("doc_id = ?", docId).Scan(&permissions).Error

	if err != nil {
		return nil, err
	}
	return permissions, nil
}

func (r *DocsRepository) UpdateDoc(doc *models.Doc) (*models.Doc, error) {
	if err := r.db.Save(doc).Error; err != nil {
		return nil, err
	}
	return doc, nil
}

func (r *DocsRepository) GetDueDocs(now time.Time) ([]models.Doc, error) {
	var docs []models.Doc

	today := now.Format("2006-01-02")
	log.Println(today)

	err := r.db.
		Table("docs").
		Where("notify_date = ? AND notif_sent = ?", today, false).
		Scan(&docs).Error
	return docs, err
}

func (r *DocsRepository) MarkNotified(docId uint) error {
	return r.db.Table("docs").Where("id = ?", docId).Update("notif_sent", true).Error
}

func (r *DocsRepository) MarkNotifCreated(docId uint) error {
	return r.db.Table("docs").Where("id = ?", docId).Update("notif_created", true).Error
}

func (r *DocsRepository) GetStatistics(docStatsDto dto.GetDocStatsDto) (*dto.DocStatsResponse, error) {
	var byStatus []dto.StatusCount

	statsQuery := r.db.Model(&models.Doc{}).
		Select("status, COUNT(*) as count")
	statsQuery = getStatsParams(docStatsDto, statsQuery)
	if err := statsQuery.
		Group("status").
		Scan(&byStatus).Error; err != nil {
		return nil, err
	}

	var byCategory []dto.CategoryCount
	categsQuery := r.db.Table("docs").
		Select(`
			docs.category_id,
			categories.name as category_name,
			categories.parent_id,
			COUNT(docs.id) as count
		`).
		Joins("LEFT JOIN categories ON categories.id = docs.category_id")
	categsQuery = getStatsParams(docStatsDto, categsQuery)
	if err := categsQuery.
		Group("docs.category_id, categories.name, categories.parent_id").
		Scan(&byCategory).Error; err != nil {
		return nil, err
	}

	var bySubCategory []dto.CategoryCount
	subCategsQuery := r.db.Table("docs").
		Select(`
			docs.sub_category_id as category_id,
			categories.name as category_name,
			categories.parent_id,
			COUNT(docs.id) as count
		`).
		Joins("LEFT JOIN categories ON categories.id = docs.sub_category_id").
		Where("docs.sub_category_id IS NOT NULL")

	subCategsQuery = getStatsParams(docStatsDto, subCategsQuery)
	if err := subCategsQuery.
		Group("docs.sub_category_id, categories.name, categories.parent_id").
		Scan(&bySubCategory).Error; err != nil {
		return nil, err
	}

	var total int64
	totalQuery := r.db.Table("docs")
	totalQuery = getStatsParams(docStatsDto, totalQuery)
	if err := totalQuery.Count(&total).Error; err != nil {
		return nil, err
	}

	stats := &dto.DocStatsResponse{
		TotalDocs:     total,
		ByStatus:      byStatus,
		ByCategory:    byCategory,
		BySubCategory: bySubCategory,
	}

	return stats, nil
	//

}

func getStatsParams(docStatsDto dto.GetDocStatsDto, query *gorm.DB) *gorm.DB {
	// Apply filters
	layout := "2006-01-02"

	if docStatsDto.DateFrom != nil {
		dateFrom, err := time.Parse(layout, *docStatsDto.DateFrom)
		if err == nil {
			if docStatsDto.DateType == "prepared" {
				query = query.Where("docs.prepared_date >= ?", dateFrom)
			} else {
				query = query.Where("docs.created_at >= ?", dateFrom)
			}
		}
	}
	if docStatsDto.DateTo != nil {
		dateTo, err := time.Parse(layout, *docStatsDto.DateTo)
		if err == nil {
			if docStatsDto.DateType == "prepared" {
				query = query.Where("docs.prepared_date <= ?", dateTo)
			} else {
				query = query.Where("docs.created_at <= ?", dateTo)
			}
		}

	}
	if docStatsDto.UserIds != nil && len(docStatsDto.UserIds) > 0 {
		query = query.Where("docs.user_id IN ?", docStatsDto.UserIds) // pass slice directly
	}
	return query
}
