package handler

import (
	"docs-notify/internal/modules/docs/service"

	"github.com/labstack/echo/v4"
)

type DocsHandler struct {
	docsService *service.DocsService
}

func NewDocsHandler(docsService *service.DocsService) *DocsHandler {
	return &DocsHandler{docsService: docsService}
}

func (h *DocsHandler) Login(c echo.Context) error {
	// Implement the login logic here
	return nil
}

func (h *DocsHandler) Create(c echo.Context) error {
	// Implement the document creation logic here
	return nil
}
