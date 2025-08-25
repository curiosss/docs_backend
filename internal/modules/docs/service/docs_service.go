package service

import (
	"docs-notify/internal/config"
	"docs-notify/internal/models"
	"docs-notify/internal/modules/docs/repository"
	fileutils "docs-notify/internal/utils/file_utils"
	"mime/multipart"

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
func (s *DocsService) CreateDoc(doc *models.Doc, file *multipart.FileHeader) error {
	// Save file locally
	filePath, err := fileutils.SaveUploadedFile(file)
	if err != nil {
		return err
	}
	// Attach to Doc model
	doc.File = filePath

	err = s.repository.CreateDoc(doc)

	return err
}
