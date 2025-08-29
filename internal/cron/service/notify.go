package service

import (
	"log"

	"docs-notify/internal/fcm/service"

	"gorm.io/gorm"
)

type NotificationService struct {
	DB         *gorm.DB
	FCMService *service.FCMService
}

// FCMSender is a wrapper around Firebase client

func (f *NotificationService) SendNotification(token string, title string, body string) error {
	// TODO: integrate Firebase SDK here
	log.Printf("Sending notification to %s: %s - %s\n", token, title, body)
	return nil
}
