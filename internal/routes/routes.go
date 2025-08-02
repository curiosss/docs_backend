package routes

import (
	"docs-notify/internal/handlers"
	"docs-notify/internal/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterUserRoutes(g *echo.Group, userHandler *handlers.UserHandler, jwtSecret string) {
	userGroup := g.Group("/users")
	userGroup.POST("/register", userHandler.Register)
	userGroup.POST("/login", userHandler.Login)
}

func RegisterDocRoutes(g *echo.Group, docHandler *handlers.DocHandler, jwtSecret string) {
	docGroup := g.Group("/docs", middleware.JWTAuth(jwtSecret))
	docGroup.POST("", docHandler.CreateDoc)
	docGroup.GET("/:id", docHandler.GetDoc)
	// Добавьте остальные роуты для документов
}

func RegisterFileRoutes(g *echo.Group, fileHandler *handlers.FileHandler, jwtSecret string) {
	fileGroup := g.Group("/files", middleware.JWTAuth(jwtSecret))
	fileGroup.POST("/upload", fileHandler.UploadFile)
}
