package service

import (
	"docs-notify/internal/config"
	"docs-notify/internal/models"
	"docs-notify/internal/modules/docs/dto"
	"docs-notify/internal/modules/docs/repository"
	"docs-notify/internal/utils/exceptions"
	fileutils "docs-notify/internal/utils/file_utils"
	"encoding/json"
	"fmt"
	"mime/multipart"
	"time"

	"gorm.io/gorm"
)

type DocsService struct {
	repository *repository.DocsRepository
	config     *config.Config
	db         *gorm.DB
}

func NewDocsService(userRepository *repository.DocsRepository, cfg *config.Config, d *gorm.DB) *DocsService {
	return &DocsService{
		repository: userRepository,
		config:     cfg,
		db:         d,
	}
}

func (s *DocsService) CreateDoc(docDto *dto.DocCreateDto, file *multipart.FileHeader) (*models.Doc, error) {
	// Save file locally
	filePath, err := fileutils.SaveUploadedFile(file)
	if err != nil {
		return nil, err
	}

	// Attach to Doc model
	layout := "2006-01-02" // Go's reference date for parsing YYYY-MM-DD

	fmt.Println(docDto.EndDate, docDto.NotifyDate)
	endDate, err := time.Parse(layout, docDto.EndDate)
	if err != nil {
		return nil, fmt.Errorf("invalid end date: %w", err)
	}
	notifyDate, err := time.Parse(layout, docDto.NotifyDate)
	if err != nil {
		return nil, fmt.Errorf("invalid notify date: %w", err)
	}
	doc := &models.Doc{
		UserId:     docDto.UserId,
		CategoryID: docDto.CategoryID,
		DocName:    docDto.DocName,
		DocNo:      docDto.DocNo,
		EndDate:    endDate,
		NotifyDate: notifyDate,
		Status:     docDto.Status,
		Permission: docDto.Permission,
		File:       filePath,
	}

	dc, err := s.repository.CreateDoc(doc)
	if err != nil {
		return nil, err
	}

	if docDto.Permissions != nil {
		var docUsers []models.DocUser

		if err := json.Unmarshal([]byte(*docDto.Permissions), &docUsers); err != nil {
			return nil, exceptions.NewResponseError(
				exceptions.ErrBadRequest,
				fmt.Errorf("invalid permissions format: %w", err),
			)
		}
		for i := range docUsers {
			docUsers[i].DocID = dc.ID
		}

		if err = s.repository.CreateDocUsers(docUsers); err != nil {
			return nil, err
		}
	}

	return dc, nil
}

func (s *DocsService) GetDocs(userId uint) ([]models.Doc, error) {
	return s.repository.GetDocs(userId)
}
