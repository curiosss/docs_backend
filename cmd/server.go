package cmd

import (
	"docs-notify/internal/config"
	"docs-notify/internal/database"

	"docs-notify/internal/utils/errorHandler"
	"docs-notify/internal/utils/validator"

	"fmt"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Server struct {
	Echo     *echo.Echo
	Config   *config.Config
	Database *gorm.DB
}

func NewServer() *Server {
	cfg := config.LoadConfig()

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
		Echo:     e,
		Config:   cfg,
		Database: db,
	}
}
