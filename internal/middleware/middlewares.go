package middleware

import (
	"docs-notify/cmd"
	"net/http"

	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
)

func RegisterMiddlewares(server *cmd.Server) {
	server.Echo.Pre(echoMiddleware.RemoveTrailingSlash())
	server.Echo.Use(
		echoMiddleware.BodyLimitWithConfig(
			echoMiddleware.BodyLimitConfig{
				Limit: "250M",
			},
		),
	)

	server.Echo.Use(
		echoMiddleware.CORSWithConfig(
			echoMiddleware.CORSConfig{
				AllowMethods: []string{
					http.MethodGet,
					http.MethodPut,
					http.MethodPatch,
					http.MethodPost,
					http.MethodDelete,
					http.MethodOptions,
				},
				AllowCredentials: true,
				AllowHeaders: []string{
					echo.HeaderOrigin,
					echo.HeaderContentType,
					echo.HeaderAccept,
					echo.HeaderAuthorization,
				},
			},
		),
	)
	server.Echo.Use(
		echoMiddleware.LoggerWithConfig(
			echoMiddleware.LoggerConfig{
				Format:           "${time_custom} | ${method} ${uri} | Status: ${status} | Error: ${error} \n",
				CustomTimeFormat: "2006/01/02 15:04:05",
			},
		),
	)
	// server.Echo.Use(echoMiddleware.RateLimiter(echoMiddleware.NewRateLimiterMemoryStore(20)))
}
