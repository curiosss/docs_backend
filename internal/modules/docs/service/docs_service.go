package service

import (
	"docs-notify/internal/config"
	"docs-notify/internal/models"
	"docs-notify/internal/modules/docs/dto"
	"docs-notify/internal/modules/docs/repository"
	fileutils "docs-notify/internal/utils/file_utils"
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
func (s *DocsService) CreateDoc(docDto *dto.DocCreateDto, file *multipart.FileHeader) error {
	// Save file locally
	filePath, err := fileutils.SaveUploadedFile(file)
	if err != nil {
		return err
	}

	// Attach to Doc model
	layout := "2006-01-02" // Go's reference date for parsing YYYY-MM-DD

	fmt.Println(docDto.EndDate, docDto.NotifyDate)
	endDate, err := time.Parse(layout, docDto.EndDate)
	if err != nil {
		return fmt.Errorf("invalid end date: %w", err)
	}
	notifyDate, err := time.Parse(layout, docDto.NotifyDate)
	if err != nil {
		return fmt.Errorf("invalid notify date: %w", err)
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
	if docDto.Permissions != nil {
		fmt.Println(docDto.Permissions)
	}

	err = s.repository.CreateDoc(doc)

	return err
}

func (s *DocsService) GetDocs(userId uint) ([]models.Doc, error) {
	return s.repository.GetDocs(userId)
}
