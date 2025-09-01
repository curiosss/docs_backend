package http

import (
	"docs-notify/cmd"
	"docs-notify/internal/middleware"

	"docs-notify/internal/modules/notifications/handler"
	"docs-notify/internal/modules/notifications/repository"
)

func InitNotificationsRouter(server *cmd.Server) {

	notifsRepo := repository.NewNotifsRepository(server.Database)
	notifsHandler := handler.NewNotifsHandler(notifsRepo)

	notifsRouter := server.Echo.Group("/api/notifications", middleware.AuthMiddleware(server))

	notifsRouter.GET("/user", notifsHandler.GetUserNotifications)
	notifsRouter.GET("/admin", notifsHandler.GetAdminNotifications, middleware.RoleMiddleware("admin"))
	// notifsRouter.POST("/create", notifsHandler.Create, middleware.RoleMiddleware("admin"))
	// notifsRouter.PUT("/update", notifsHandler.Update, middleware.RoleMiddleware("admin"))
	// notifsRouter.GET("/delete", notifsHandler.Delete, middleware.RoleMiddleware("admin"))
}
