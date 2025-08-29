package main

import (
	"context"
	"docs-notify/cmd"
	"docs-notify/internal"
	cronjob "docs-notify/internal/cron/job"
	"docs-notify/internal/middleware"
	"docs-notify/internal/modules/docs/repository"
	usersRepo "docs-notify/internal/modules/users/repository"

	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-co-op/gocron"
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

	server := cmd.NewServer()

	middleware.RegisterMiddlewares(server)

	// Статические файлы для загрузок
	server.Echo.Static("/uploads/docs", "uploads/docs")

	// Swagger
	server.Echo.GET("/docs/*", echoSwagger.WrapHandler)

	internal.InitRouters(server)

	docsRepo := repository.NewDocsRepository(server.Database)
	usersRepo := usersRepo.NewUserRepository(server.Database)
	notifyCron := cronjob.NewNotifyCron(docsRepo, usersRepo, server.FCMService)

	// Запуск сервера в горутине
	go func() {
		if err := server.Echo.Start(":" + server.Config.AppPort); err != nil && err != http.ErrServerClosed {
			server.Echo.Logger.Fatal("shutting down the server")
		}
	}()

	// Setup scheduler (every 1 hour, only between 08:00–22:00 UTC)
	s := gocron.NewScheduler(time.Local)
	s.Every(5).Minute().Do(notifyCron.Run)

	// Run scheduler async
	s.StartAsync()

	// Ожидание сигнала прерывания для graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := server.Echo.Shutdown(ctx); err != nil {
		server.Echo.Logger.Fatal(err)
	}

	// Stop cron safely
	s.Stop()
}
