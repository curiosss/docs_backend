package service

import (
	"docs-notify/internal/config"
	"docs-notify/internal/models"
	"docs-notify/internal/modules/docs/repository"

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
func (s *DocsService) CreateDocWithFiles(doc *models.Doc, files []models.File) error {
	// 1. Create the doc
	if err := s.repository.CreateDoc(doc); err != nil {
		return err
	}

	// 2. Attach docID to each file
	for i := range files {
		files[i].DocID = doc.ID
	}

	// 3. Save files
	if err := s.repository.CreateFiles(files); err != nil {
		return err
	}

	return nil
}
