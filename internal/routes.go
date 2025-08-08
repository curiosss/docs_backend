package internal

import (
	"docs-notify/cmd"
	usersRouter "docs-notify/internal/modules/users/http"
)

func InitRouters(server *cmd.Server) {
	usersRouter.InitUsersRouter(server)
}
