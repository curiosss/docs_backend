package http

import (
	"docs-notify/cmd"
	"docs-notify/internal/middleware"
	"docs-notify/internal/modules/categories/handler"
	"docs-notify/internal/modules/categories/repository"
)

func InitCategoriesRouter(server *cmd.Server) {

	categoriesRepo := repository.NewCategRepository(server.Database)
	categoriesHandler := handler.NewCategHandler(categoriesRepo)

	categRouter := server.Echo.Group("/api/categories", middleware.AuthMiddleware(server))

	categRouter.GET("/all", categoriesHandler.GetAll)
	categRouter.POST("/create", categoriesHandler.Create, middleware.RoleMiddleware("admin"))
	categRouter.PUT("/update", categoriesHandler.Update, middleware.RoleMiddleware("admin"))
	categRouter.GET("/delete", categoriesHandler.Delete, middleware.RoleMiddleware("admin"))
}
