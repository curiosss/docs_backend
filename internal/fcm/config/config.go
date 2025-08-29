package app

import (
	"context"
	"log"

	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

func InitFirebase() (*firebase.App, error) {
	opt := option.WithCredentialsFile("internal/fcm/config/docs-6f87c-firebase-adminsdk-fbsvc-397fefb08c.json") // downloaded from Firebase console
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Printf("error initializing firebase app: %v", err)
	}
	return app, err
}
