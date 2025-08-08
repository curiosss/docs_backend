package jobs

// import (
// 	"log"

// 	"docs-notify/internal/services"
// )

// type NotifyJob struct {
// 	docService    services.DocService
// 	notifyService services.NotifyService
// }

// func NewNotifyJob(ds services.DocService, ns services.NotifyService) *NotifyJob {
// 	return &NotifyJob{
// 		docService:    ds,
// 		notifyService: ns,
// 	}
// }

// func (j *NotifyJob) Run() {
// 	log.Println("Running notify job...")

// 	docs, err := j.docService.GetDocsForNotification()
// 	if err != nil {
// 		log.Printf("Error getting docs for notification: %v", err)
// 		return
// 	}

// 	if len(docs) == 0 {
// 		log.Println("No documents to notify about.")
// 		return
// 	}

// 	for _, doc := range docs {
// 		log.Printf("Processing doc ID: %d", doc.ID)

// 		// Здесь должна быть логика получения токенов пользователей, подписанных на документ
// 		// userTokens := j.docService.GetUserTokensForDoc(doc.ID)
// 		userTokens := []string{"dummy-token-for-testing"} // Заглушка

// 		err := j.notifyService.Send(doc, userTokens)
// 		if err != nil {
// 			log.Printf("Error sending notification for doc ID %d: %v", doc.ID, err)
// 			continue
// 		}

// 		err = j.docService.MarkAsNotified(doc.ID)
// 		if err != nil {
// 			log.Printf("Error marking doc ID %d as notified: %v", doc.ID, err)
// 		}
// 	}

// 	log.Println("Notify job finished.")
// }
