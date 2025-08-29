package service

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
	client, err := app.Messaging(context.Background())
	if err != nil {
		return nil, fmt.Errorf("error getting messaging client: %w", err)
	}

	// msg := &messaging.Message{
	// 	Token: "fbYJzkwyTvGsY9XrjYZ1D-:APA91bGgR_FqU7hw60s2LjuPgo1dSpRmz9WnEmnJaJbaBAvV0hvAy2tRaPvIchxutKSxRmPYxS4w5QHEtaPi2kGc7qG_wmhDH2ac8hzxgdeOb4cPIYtRAAk",
	// 	Notification: &messaging.Notification{
	// 		Title: "Big Announcement 🎉",
	// 		Body:  "This goes to everyone!",
	// 	},
	// 	// Topic: "all",
	// }
	// res, err := client.Send(context.Background(), msg)
	// if err != nil {
	// 	log.Fatalf("error sending: %v", err)
	// 	return nil, err
	// }

	// log.Println("FCM client initialized")
	// fmt.Printf("Successfully sent message: %s\n", res)

	return &FCMService{client: client}, nil
}

func (s *FCMService) SendMessage(title string, body string, fcmToken string) error {
	// Send to ALL (via topic "all")
	log.Println(fcmToken)
	msg := &messaging.Message{
		Token: fcmToken,
		Notification: &messaging.Notification{
			Title: title,
			Body:  body,
		},
	}

	resp, err := s.client.Send(context.Background(), msg)
	if err != nil {
		log.Printf("error sending message: %v", err)
		return err
	}
	log.Printf("Successfully sent message: %s", resp)
	return nil
}
