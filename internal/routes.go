package internal

import (
	"docs-notify/cmd"
	categoriesRouter "docs-notify/internal/modules/categories/http"
	docsRouter "docs-notify/internal/modules/docs/http"
	notifsRouter "docs-notify/internal/modules/notifications/http"
	usersRouter "docs-notify/internal/modules/users/http"
)

func InitRouters(server *cmd.Server) {
	usersRouter.InitUsersRouter(server)
	docsRouter.InitDocsRouter(server)
	categoriesRouter.InitCategoriesRouter(server)
	notifsRouter.InitNotificationsRouter(server)
}
