package app

import (
	"context"
	"log"

	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

func InitFirebase() (*firebase.App, error) {
	opt := option.WithCredentialsFile("internal/fcm/config/fcm_credentials.json")
	// opt := option.WithCredentialsJSON([]byte(`
	// {
	//   "type": "service_account",
	//   "project_id": "docs-6f87c",
	//   "private_key_id": "0a273b15092e4a0c811816bfb4816a9068c788f5",
	//   "private_key": "-----BEGIN PRIVATE KEY-----\nMIIEvAIBADANBgkqhkiG9w0BAQEFAASCBKYwggSiAgEAAoIBAQC2GLZdlQArpjcH\nC569+bQk5j/86VOpDbxxcInnAo965yqle4meyd+PwcydrzPl4UUY6b5E107DOIS8\n7JrEsjYTecDWplHh8UxoRox+J07lxGLY3ZEnsacnG4en8HQ8RQyiwHvyM5U6L88p\namXnw/cTcJIF6cBqQIR/40faaJ/ULHhqr+tvubyyehDTdG4Rh1oLb+6pD3oqQjKs\njRuxn85ht3y8Bbvp2SZmXRC80W2NfXMvFC3YaIFI+tGBiELZudR4evThT7GwTiH0\nveituYBd4yYD0tY2pd+tEcFWzeHfBuY6ejF0FQm65S6UiJ8nLGvvXsQpwDgvLTYf\npk70ujVFAgMBAAECggEAJ1xUHwpUDRtSja1PVNUiwU8byblNoh8b+pFO3aZKCVDq\ngPahGrecIWDMr9DtMOVuoCH2RL8dgUk4N/YUxiuXUMJikoNE88fKskd9ms84WKYI\nj8Pk2mWmMefbfMuW9GgggRWGNSY9PWSULOmYuW48e/E7Pxf1xYHIojQoledDlQLi\nU9UmZHuI8XoP2tfp8yuFHOJTcXl0d5+h9To77gPdUk0Py8XALwDAzLUmtpWBdelZ\nbLpPUHGiEunqAV+cRg5sH5rzcz9RHbYHjK6hQw+WXuOssGaIo6qrstmXeQn11IAP\nYp+SiDYSI5IAgQ2wHXaQi1xLezD8kdB5MKYRwjWfiQKBgQDjqpYwvI7YTGhNwU9C\nZ/dKB6/gifRyhmmfMATlcVYRKHh4peQGnyFSVnl2QvzMZsKISIKs4Ai584o3SKr+\nxujc85lqUgd2qNoHJ6DtIcJxVsHSlzJhme0r2pwSMfze3+3AGr4ix7YqvTxASIrQ\nW4nPQVbwM3HxjEUXbVOyYm8aCwKBgQDMwkdNKE7UHbq0tK8T0xm5xTuxq/AqWpUk\nHV1KFsy7oV3qlzTBmlwscO7/t4XD0qWtdNTflCIASPiE6SPb/O/wJTpAwJpy9a/K\nOdPnGG/lBW4+25A1ZQmsh6giP1sGkPkOXaFKa+F8LsaH0A3ViIqvcHbOuw8jfbJk\nuJ/8dBbP7wKBgFtL6CipJLtWgKlVsNwXZxJX3M61Y8KdZjPBBOWhunrs+Mqg8704\nCRvEs7aaDFhHiREvyr9apAU1xaJ/0JqU14LraQU62eVatvwRhzYwyJG80cMKgNik\n6ngglV+yjg4uTGAyGTdHUST4d/XrYUdGvg/PyvZOGw5bSsWnQN4THSybAoGAAc5M\n1q9eUpyYgvN8/83C0lKc/iooCheWbSUdJ4Qf9h+sNl9zBaoY2gN8+CBkO5/l+iun\nnPkve5UpK/LqcAxBCXsqkluggRcNn9j2t3kNs5VirYc+NFpZxX3Ey9iHMv2gVLIa\ntA9Tg8bd1WDOXm2/22BAi/42WffH1P+T2aQkd10CgYAuoe/+nFKmpzBzntx9ncJ+\nM9+mmszvIgDpk+2H206h9JONd8fG8l/AUA6y22FAlVectjFmBlbbYKOqhoGC5gZw\nBAtBUpLuvZ7L/0rZY+qyDZl8Xq/G1MLprip3NfoyHY0RdZNz0eCCkuCzfP9F5VD3\nxK6lcQf6e3ACmE2D4Ld10Q==\n-----END PRIVATE KEY-----\n",
	//   "client_email": "firebase-adminsdk-fbsvc@docs-6f87c.iam.gserviceaccount.com",
	//   "client_id": "108526452350016007488",
	//   "auth_uri": "https://accounts.google.com/o/oauth2/auth",
	//   "token_uri": "https://oauth2.googleapis.com/token",
	//   "auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
	//   "client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/firebase-adminsdk-fbsvc%40docs-6f87c.iam.gserviceaccount.com",
	//   "universe_domain": "googleapis.com"
	// }
	// 	`))
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Printf("error initializing firebase app: %v", err)
	}
	return app, err
}
