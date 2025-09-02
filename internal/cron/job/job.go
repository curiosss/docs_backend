package cronjob

import (
	"docs-notify/internal/fcm/service"
	"docs-notify/internal/models"
	docsRepo "docs-notify/internal/modules/docs/repository"
	notifsRepo "docs-notify/internal/modules/notifications/repository"
	"docs-notify/internal/modules/users/dto"
	usersRepo "docs-notify/internal/modules/users/repository"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"
)

type NotifyCron struct {
	docsRepo   *docsRepo.DocsRepository
	usersRepo  *usersRepo.UserRepository
	notifsRepo *notifsRepo.NotifsRepository
	fcm        *service.FCMService
}

func NewNotifyCron(
	repo *docsRepo.DocsRepository,
	userRepo *usersRepo.UserRepository,
	notifRepo *notifsRepo.NotifsRepository,
	fcm *service.FCMService,
) *NotifyCron {
	return &NotifyCron{
		docsRepo:   repo,
		usersRepo:  userRepo,
		notifsRepo: notifRepo,
		fcm:        fcm,
	}
}

func (c *NotifyCron) Run() {
	nowUtc := time.Now().UTC()
	now := time.Now()
	// Only run between 08:00 and 18:00
	hour := nowUtc.Hour()
	if hour < 4 || hour > 18 {
		log.Printf("⏸ Skipping notifications, outside time window (%d:00)", hour)
		return
	}

	docs, err := c.docsRepo.GetDueDocs(now)
	if err != nil {
		log.Printf("❌ error fetching docs: %v", err)
		return
	}

	users, err := c.usersRepo.GetAllNotif()
	if err != nil {
		log.Printf("❌ error fetching users: %v", err)
		return
	}

	for _, doc := range docs {
		if doc.NotifSent {
			continue
		}
		if doc.Permission != nil && *doc.Permission > 0 {

			for _, user := range users {
				if err = c.SendNotification(&doc, &user); err != nil {
					log.Println(err)
					continue
				}
			}

		} else {
			docUsers, err := c.docsRepo.GetDocPermissions(doc.ID)
			if err != nil {
				log.Printf("error fetching doc users for doc %d: %v", doc.ID, err)
				continue
			}

			for _, du := range docUsers {
				for _, user := range users {
					if user.ID == du.UserID {
						if err = c.SendNotification(&doc, &user); err != nil {
							log.Println(err)
							continue
						}
					}
				}

			}
		}

	}
}

func (c *NotifyCron) SendNotification(doc *models.Doc, user *dto.UserNotifDto) error {
	notif := &models.Notification{
		DocID:  doc.ID,
		UserID: user.ID,
		Title:  "Dokument duýdurmaly wagty geldi",
		Body:   "Dokument " + doc.DocName + " , gutaryan wagty " + doc.NotifyDate.Format("2006-01-02"),
		IsSeen: false,
	}

	if !doc.NotifCreated {
		updated, err := c.notifsRepo.Add(notif)
		if err != nil {
			return errors.New(fmt.Sprintf("failed to add notifications: %v", err))
		}
		notif = updated
		_ = c.docsRepo.MarkNotifCreated(doc.ID)
	}

	if user.FcmToken == "" {
		return errors.New("fcm token is empty")
	}

	err := c.fcm.SendMessage(
		notif.Title,
		notif.Body,
		user.FcmToken,
	)
	if err != nil {

		if strings.Contains(strings.ToLower(err.Error()), "requested entity was not found") {
			c.usersRepo.RemoveFcmToken(user.ID)
		}
		return errors.New(fmt.Sprintf("failed to send notification for doc %d: %v", doc.ID, err))
	}

	_ = c.docsRepo.MarkNotified(doc.ID)
	log.Printf("Notification sent for doc %d", doc.ID)
	return nil
}
