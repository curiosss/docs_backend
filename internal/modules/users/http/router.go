package http

import (
	"docs-notify/cmd"
	"docs-notify/internal/middleware"
	"docs-notify/internal/modules/users/handler"
	"docs-notify/internal/modules/users/repository"
	"docs-notify/internal/modules/users/service"
)

func InitUsersRouter(server *cmd.Server) {

	userRepository := repository.NewUserRepository(server.Database)
	userService := service.NewUserService(userRepository, server.Config)
	userHandler := handler.NewUserHandler(userService)

	server.Echo.POST("/api/users/login", userHandler.Login)

	userRouter := server.Echo.Group("/api/users", middleware.AuthMiddleware(server))

	userRouter.GET("/all", userHandler.GetAll)
	userRouter.PUT("/change-username", userHandler.ChangeUsername)
	userRouter.PUT("/change-password", userHandler.ChangePassword)
	userRouter.POST("/create", userHandler.CreateUser, middleware.RoleMiddleware("admin"))
	userRouter.PUT("/update", userHandler.UpdateUser, middleware.RoleMiddleware("admin"))
	userRouter.GET("/delete", userHandler.DeleteUser, middleware.RoleMiddleware("admin"))
}
