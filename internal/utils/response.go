package utils

import (
	"github.com/labstack/echo/v4"
)

type SuccessResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type ErrorResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func SendSuccess(c echo.Context, statusCode int, message string, data interface{}) error {
	return c.JSON(statusCode, SuccessResponse{
		Success: true,
		Message: message,
		Data:    data,
	})
}

func SendError(c echo.Context, statusCode int, message interface{}) error {
	msgStr := "An error occurred"
	if m, ok := message.(string); ok {
		msgStr = m
	}
	if m, ok := message.(error); ok {
		msgStr = m.Error()
	}

	return c.JSON(statusCode, ErrorResponse{
		Success: false,
		Message: msgStr,
	})
}
