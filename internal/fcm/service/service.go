package fcm

import (
	"context"
	"fmt"
	"log"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"
)

type FCMService struct {
	client *messaging.Client
}

func NewFCMService(app *firebase.App) (*FCMService, error) {
	ctx := context.Background()
	client, err := app.Messaging(ctx)
	if err != nil {
		return nil, fmt.Errorf("error getting messaging client: %w", err)
	}

	// msg := &messaging.Message{
	// 	Topic: "all",
	// 	Notification: &messaging.Notification{
	// 		Title: "Big Announcement 🎉",
	// 		Body:  "This goes to everyone!",
	// 	},
	// }
	// res, err := client.SendAll(ctx, []*messaging.Message{msg})
	// if err != nil {
	// 	log.Fatalf("error sending: %v", err)
	// 	return nil, err
	// }
	// log.Println("FCM client initialized")
	// fmt.Printf("Successfully sent message: %s\n", res)
	return &FCMService{client: client}, nil
}

func (s *FCMService) SendNotification(ctx context.Context, token, title, body string, data map[string]string) (string, error) {
	message := &messaging.Message{
		Token: token,
		Notification: &messaging.Notification{
			Title: title,
			Body:  body,
		},
		Data: data,
	}

	// Send message
	response, err := s.client.Send(ctx, message)
	if err != nil {
		return "", fmt.Errorf("error sending message: %w", err)
	}
	fmt.Printf("Successfully sent message: %s\n", response)

	return response, nil
}

func (s *FCMService) SendAll() error {
	// Send to ALL (via topic "all")
	msg := &messaging.Message{
		Topic: "all",
		Notification: &messaging.Notification{
			Title: "Big Announcement ",
			Body:  "This goes to everyone!",
		},
	}

	fmt.Printf("Sending message with client", s.client)

	resp, err := s.client.SendAll(context.Background(), []*messaging.Message{msg})
	if err != nil {
		log.Fatalf("error sending: %v", err)
		return err
	}
	log.Printf("Successfully sent: %s", resp)

	return nil
}
