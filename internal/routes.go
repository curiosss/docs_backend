package internal

import (
	"docs-notify/cmd"
	docsRouter "docs-notify/internal/modules/docs/http"
	usersRouter "docs-notify/internal/modules/users/http"
)

func InitRouters(server *cmd.Server) {
	usersRouter.InitUsersRouter(server)
	docsRouter.InitDocsRouter(server)
}
