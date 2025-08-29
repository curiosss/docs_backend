package cronjob

import (
	"docs-notify/internal/fcm/service"
	docsRepo "docs-notify/internal/modules/docs/repository"
	usersRepo "docs-notify/internal/modules/users/repository"
	"log"
	"time"
)

type NotifyCron struct {
	docsRepo  *docsRepo.DocsRepository
	usersRepo *usersRepo.UserRepository
	fcm       *service.FCMService
}

func NewNotifyCron(repo *docsRepo.DocsRepository, userRepo *usersRepo.UserRepository, fcm *service.FCMService) *NotifyCron {
	return &NotifyCron{docsRepo: repo, usersRepo: userRepo, fcm: fcm}
}

func (c *NotifyCron) Run() {
	nowUtc := time.Now().UTC()
	now := time.Now()
	// Only run between 08:00 and 18:00
	hour := nowUtc.Hour()
	if hour < 13 || hour > 23 {
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
		if doc.Permission != nil && *doc.Permission > 0 {

			for _, user := range users {
				if user.FcmToken == "" {
					continue
				}
				err := c.fcm.SendMessage(
					"Dokument duýdurmaly wagty geldi",
					"Dokument "+doc.DocName+" , gutaryan wagty "+doc.NotifyDate.Format("2006-01-02"),
					user.FcmToken)

				if err != nil {
					log.Printf("❌ failed to send notification for doc %d: %v", doc.ID, err)
					continue
				}
				_ = c.docsRepo.MarkNotified(doc.ID)
				log.Printf("✅ Notification sent for doc %d", doc.ID)
			}

		} else {
			docUsers, err := c.docsRepo.GetDocPermissions(doc.ID)
			if err != nil {
				log.Printf("❌ error fetching doc users for doc %d: %v", doc.ID, err)
				continue
			}

			for _, du := range docUsers {

				for _, user := range users {
					if user.ID == du.UserID {
						if user.FcmToken == "" {
							continue
						}
						err := c.fcm.SendMessage(
							"Dokument duýdurmaly wagty geldi",
							"Dokument "+doc.DocName+" , gutaryan wagty "+doc.NotifyDate.Format("2006-01-02"),
							user.FcmToken)

						if err != nil {
							log.Printf("❌ failed to send notification for doc %d: %v", doc.ID, err)
							continue
						}
						_ = c.docsRepo.MarkNotified(doc.ID)
						log.Printf("✅ Notification sent for doc %d", doc.ID)
					}
				}

			}
		}

	}
}

// func sendNotifications(doc *models.Doc, users []dto.UserNotifDto) {
// 	for _, user := range users {
// 		if user.FcmToken == "" {
// 			continue
// 		}
// 		err := c.fcm.SendMessage(
// 			"Dokument duýdurmaly wagty geldi",
// 			"Dokument "+doc.DocName+" , gutaryan wagty "+doc.NotifyDate.Format("2006-01-02 15:04"),
// 			user.FcmToken)

// 		if err != nil {
// 			log.Printf("❌ failed to send notification for doc %d: %v", doc.ID, err)
// 			continue
// 		}
// 		_ = c.docsRepo.MarkNotified(doc.ID)
// 		log.Printf("✅ Notification sent for doc %d", doc.ID)
// 	}
// }
