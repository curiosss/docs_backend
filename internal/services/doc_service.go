package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"docs-notify/internal/dto"
	"docs-notify/internal/models"
	"docs-notify/internal/repositories"
)

// DocService
type DocService interface {
	CreateDoc(req *dto.CreateDocRequest, authorID uint) (*dto.DocResponse, error)
	GetDoc(id uint) (*dto.DocResponse, error)
	GetDocsForNotification() ([]models.Doc, error)
	MarkAsNotified(docID uint) error
	AddFileToDoc(docID uint, filename, filepath, url string) (*models.File, error)
}

type docService struct {
	repo repositories.DocRepository
}

func NewDocService(repo repositories.DocRepository) DocService {
	return &docService{repo}
}

func (s *docService) CreateDoc(req *dto.CreateDocRequest, authorID uint) (*dto.DocResponse, error) {
	var notifyDate time.Time
	if req.NotifyDate != "" {
		parsedTime, err := time.Parse(time.RFC3339, req.NotifyDate)
		if err != nil {
			return nil, fmt.Errorf("invalid date format: %w", err)
		}
		notifyDate = parsedTime
	}

	doc := &models.Doc{
		Title:      req.Title,
		Content:    req.Content,
		AuthorID:   authorID,
		CategoryID: req.CategoryID,
		NotifyDate: notifyDate,
	}

	if err := s.repo.Create(doc); err != nil {
		return nil, err
	}

	return &dto.DocResponse{
		ID:         doc.ID,
		Title:      doc.Title,
		Content:    doc.Content,
		AuthorID:   doc.AuthorID,
		CategoryID: doc.CategoryID,
		NotifyDate: doc.NotifyDate.Format(time.RFC3339),
	}, nil
}

func (s *docService) GetDoc(id uint) (*dto.DocResponse, error) {
	doc, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	return &dto.DocResponse{
		ID:         doc.ID,
		Title:      doc.Title,
		Content:    doc.Content,
		AuthorID:   doc.AuthorID,
		CategoryID: doc.CategoryID,
		NotifyDate: doc.NotifyDate.Format(time.RFC3339),
	}, nil
}

func (s *docService) GetDocsForNotification() ([]models.Doc, error) {
	return s.repo.FindForNotification()
}

func (s *docService) MarkAsNotified(docID uint) error {
	return s.repo.UpdateNotifSent(docID)
}

func (s *docService) AddFileToDoc(docID uint, filename, filepath, url string) (*models.File, error) {
	file := &models.File{
		DocID:    docID,
		Filename: filename,
		Filepath: filepath,
		URL:      url,
	}
	err := s.repo.AddFile(file)
	if err != nil {
		return nil, err
	}
	return file, nil
}

// NotifyService
type NotifyService interface {
	Send(doc models.Doc, userTokens []string) error
}

type notifyService struct {
	fcmServerKey string
}

func NewNotifyService(fcmKey string) NotifyService {
	return &notifyService{fcmServerKey: fcmKey}
}

func (s *notifyService) Send(doc models.Doc, userTokens []string) error {
	if len(userTokens) == 0 {
		log.Println("No tokens to send notification to.")
		return nil
	}

	fcmURL := "https://fcm.googleapis.com/fcm/send"

	notificationPayload := map[string]interface{}{
		"registration_ids": userTokens,
		"notification": map[string]string{
			"title": "Новый документ: " + doc.Title,
			"body":  "Появился новый документ, который может вас заинтересовать.",
		},
		"data": map[string]string{
			"doc_id": fmt.Sprintf("%d", doc.ID),
		},
	}

	payloadBytes, err := json.Marshal(notificationPayload)
	if err != nil {
		return fmt.Errorf("failed to marshal FCM payload: %w", err)
	}

	req, err := http.NewRequest("POST", fcmURL, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return fmt.Errorf("failed to create FCM request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "key="+s.fcmServerKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send FCM request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		// Здесь можно добавить более детальную обработку ответа от FCM
		return fmt.Errorf("FCM request failed with status: %s", resp.Status)
	}

	log.Printf("Successfully sent notification for doc ID %d", doc.ID)
	return nil
}
