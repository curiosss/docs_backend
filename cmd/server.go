package cmd

import (
	"docs-notify/internal/config"
	"docs-notify/internal/database"
	fcm "docs-notify/internal/fcm/config"
	"docs-notify/internal/fcm/service"
	"log"

	"docs-notify/internal/utils/errorHandler"
	"docs-notify/internal/utils/validator"

	"fmt"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Server struct {
	Echo       *echo.Echo
	Config     *config.Config
	Database   *gorm.DB
	FCMService *service.FCMService
}

func NewServer() *Server {
	cfg := config.LoadConfig()

	var fcmService *service.FCMService
	fcmApp, err := fcm.InitFirebase()
	if err == nil {
		fcmService, err = service.NewFCMService(fcmApp)
		if err != nil {
			log.Printf("Failed to initialize FCM service: %v", err)
		}
	}
	log.Println("Firebase Service initialized")

	// Подключение к БД
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort)

	db, err := database.Connect(dsn, cfg)
	if err != nil {
		panic(err)
	}

	e := echo.New()
	e.Validator = validator.NewValidator()
	e.HTTPErrorHandler = errorHandler.ResponseHTTPErrorHandler

	// Middleware
	// e.Use(echoMiddleware.Logger())
	// e.Use(echoMiddleware.Recover())
	// e.Use(middleware.ErrorHandler)

	return &Server{
		Echo:       e,
		Config:     cfg,
		Database:   db,
		FCMService: fcmService,
	}
}
