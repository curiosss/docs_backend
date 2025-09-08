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
		UserId:        docDto.UserId,
		CategoryID:    docDto.CategoryID,
		SubCategoryId: docDto.SubCategoryId,
		DocName:       docDto.DocName,
		DocNo:         docDto.DocNo,
		EndDate:       endDate,
		NotifyDate:    notifyDate,
		Status:        docDto.Status,
		Permission:    docDto.Permission,
		File:          filePath,
	}

	fmt.Println(docDto)

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

func (s *DocsService) GetDocs(getDocsDto dto.GetDocsDto) (*dto.DocsResponseDto, error) {
	return s.repository.GetDocsForUser(getDocsDto)
}

func (s *DocsService) GetDocById(docId uint) (*dto.DocResponse, error) {
	return s.repository.GetDocById(docId)
}

func (s *DocsService) DeleteDoc(docId uint) error {

	doc, err := s.repository.GetByID(docId)
	if err != nil {
		return exceptions.NewResponseError(exceptions.ErrBadRequest, err)
	}

	// Delete file from storage
	if err := fileutils.DeleteFile(doc.File); err != nil {
		return exceptions.NewResponseError(exceptions.ErrInternalServerError, err)
	}

	// Delete doc users associations
	if err := s.repository.DeleteDocUsersByDocID(docId); err != nil {
		return exceptions.NewResponseError(exceptions.ErrInternalServerError, err)
	}

	// Delete doc notifications
	if err := s.repository.DeleteDocNotifications(docId); err != nil {
		return exceptions.NewResponseError(exceptions.ErrInternalServerError, err)
	}

	err = s.repository.Delete(docId)
	return err
}

func (s *DocsService) GetDocPermissions(docId uint) ([]models.DocUser, error) {
	return s.repository.GetDocPermissions(docId)
}

func (s *DocsService) UpdateDoc(docDto *dto.DocUpdateDto, file *multipart.FileHeader) (*models.Doc, error) {
	// Fetch existing doc
	existingDoc, err := s.repository.GetByID(docDto.Id)
	if err != nil {
		return nil, exceptions.NewResponseError(exceptions.ErrBadRequest, fmt.Errorf("Dokument tapylmady: %w", err))
	}

	// Update fields
	layout := "2006-01-02" // Go's reference date for parsing YYYY-MM-DD

	endDate, err := time.Parse(layout, docDto.EndDate)
	if err != nil {
		return nil, fmt.Errorf("invalid end date: %w", err)
	}
	notifyDate, err := time.Parse(layout, docDto.NotifyDate)
	if err != nil {
		return nil, fmt.Errorf("invalid notify date: %w", err)
	}

	if existingDoc.NotifyDate != notifyDate {
		existingDoc.NotifSent = false
		existingDoc.NotifCreated = false
	}

	existingDoc.CategoryID = docDto.CategoryID
	existingDoc.SubCategoryId = docDto.SubCategoryId
	existingDoc.DocName = docDto.DocName
	existingDoc.DocNo = docDto.DocNo
	existingDoc.EndDate = endDate
	existingDoc.NotifyDate = notifyDate
	existingDoc.Status = docDto.Status
	if docDto.Status == "prepared" {
		now := time.Now()
		existingDoc.PreparedDate = &now
	}
	existingDoc.Permission = docDto.Permission

	if err = s.repository.DeleteDocUsersByDocID(existingDoc.ID); err != nil {
		return nil, err
	}

	// If new file is provided, replace the old one
	if file != nil {
		// Delete old file
		if err := fileutils.DeleteFile(existingDoc.File); err != nil {
			return nil, exceptions.NewResponseError(exceptions.ErrInternalServerError, err)
		}
		// Save new file
		filePath, err := fileutils.SaveUploadedFile(file)
		if err != nil {
			return nil, exceptions.NewResponseError(exceptions.ErrInternalServerError, err)
		}
		existingDoc.File = filePath
	}

	// Update permissions if provided
	if docDto.Permissions != nil {
		var docUsers []models.DocUser

		if err := json.Unmarshal([]byte(*docDto.Permissions), &docUsers); err != nil {
			return nil, exceptions.NewResponseError(
				exceptions.ErrBadRequest,
				fmt.Errorf("invalid permissions format: %w", err),
			)
		}
		for i := range docUsers {
			docUsers[i].DocID = existingDoc.ID
		}
		if err = s.repository.CreateDocUsers(docUsers); err != nil {
			return nil, err
		}

	}
	// Update doc in DB
	updatedDoc, err := s.repository.UpdateDoc(existingDoc)
	if err != nil {
		return nil, exceptions.NewResponseError(exceptions.ErrInternalServerError, err)
	}
	return updatedDoc, nil
}

func (s *DocsService) GetStatistics(docStatsDto dto.GetDocStatsDto) (*dto.DocStatsResponse, error) {
	return s.repository.GetStatistics(docStatsDto)
}
