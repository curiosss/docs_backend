package http

import (
	"docs-notify/cmd"
	"docs-notify/internal/middleware"
	"docs-notify/internal/modules/docs/handler"
	"docs-notify/internal/modules/docs/repository"
	"docs-notify/internal/modules/docs/service"
)

func InitDocsRouter(server *cmd.Server) {

	docsRepository := repository.NewDocsRepository(server.Database)
	docsService := service.NewDocsService(docsRepository, server.Config, server.Database)
	docsHandler := handler.NewDocsHandler(docsService)

	// server.Echo.POST("/api/users/login", docsHandler.Login)

	docsRouter := server.Echo.Group("/api/docs", middleware.AuthMiddleware(server))
	docsRouter.POST("/create", docsHandler.Create, middleware.RoleMiddleware("operator"))
}
