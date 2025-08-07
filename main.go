package main

import (
	"context"
	"docs-notify/internal/middleware"
	"docs-notify/internal/server"
	"net/http"
	"os"
	"os/signal"
	"time"

	// _ "docs-notify/docs" // load generated docs

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

	server := server.NewServer()

	middleware.RegisterMiddlewares(server)

	// Статические файлы для загрузок
	server.Echo.Static("/static", "uploads")

	// Swagger
	server.Echo.GET("/docs/*", echoSwagger.WrapHandler)

	// Запуск сервера в горутине
	go func() {
		if err := server.Echo.Start(":" + server.Config.AppPort); err != nil && err != http.ErrServerClosed {
			server.Echo.Logger.Fatal("shutting down the server")
		}
	}()

	// Ожидание сигнала прерывания для graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := server.Echo.Shutdown(ctx); err != nil {
		server.Echo.Logger.Fatal(err)
	}
}
