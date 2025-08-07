package middleware

import (
	"net/http"

	"docs-notify/internal/utils"

	"github.com/labstack/echo/v4"
)

func ErrorHandler(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := next(c)
		if err != nil {
			if he, ok := err.(*echo.HTTPError); ok {
				return utils.SendError(c, he.Code, he.Message)
			}
			// Для других типов ошибок можно добавить обработку
			return utils.SendError(c, http.StatusInternalServerError, "Internal Server Error")
		}
		return nil
	}
}
