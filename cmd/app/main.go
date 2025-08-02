package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"docs-notify/internal/config"
	"docs-notify/internal/handlers"
	"docs-notify/internal/middleware"
	"docs-notify/internal/models"
	"docs-notify/internal/repositories"
	"docs-notify/internal/routes"
	"docs-notify/internal/services"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	_ "docs-notify/docs" // load generated docs

	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Docs Notify API
// @version 1.0
// @description This is a sample server for a documentation notification service.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	// Загрузка .env файла (для локальной разработки)
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Конфигурация
	cfg := config.LoadConfig()

	// Подключение к БД
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// Миграция схемы
	db.AutoMigrate(&models.User{}, &models.Category{}, &models.Doc{}, &models.DocUser{}, &models.Action{}, &models.File{})

	// Инициализация Echo
	e := echo.New()

	// Middleware
	e.Use(echoMiddleware.Logger())
	e.Use(echoMiddleware.Recover())
	e.Use(middleware.ErrorHandler)

	// Статические файлы для загрузок
	e.Static("/static", "uploads")

	// Swagger
	e.GET("/docs/*", echoSwagger.WrapHandler)

	// Инициализация зависимостей (DI)
	// Repositories
	userRepo := repositories.NewUserRepository(db)
	docRepo := repositories.NewDocRepository(db)

	// Services
	userService := services.NewUserService(userRepo, cfg.JWTSecret)
	docService := services.NewDocService(docRepo)
	notifyService := services.NewNotifyService(cfg.FCMServerKey)

	// Handlers
	userHandler := handlers.NewUserHandler(userService)
	docHandler := handlers.NewDocHandler(docService)
	fileHandler := handlers.NewFileHandler(docService) // Используем DocService для связи файла с документом

	// Роуты
	apiGroup := e.Group("/api/v1")
	routes.RegisterUserRoutes(apiGroup, userHandler, cfg.JWTSecret)
	routes.RegisterDocRoutes(apiGroup, docHandler, cfg.JWTSecret)
	routes.RegisterFileRoutes(apiGroup, fileHandler, cfg.JWTSecret)

	// Запуск сервера в горутине
	go func() {
		if err := e.Start(":" + cfg.AppPort); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	// Ожидание сигнала прерывания для graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
