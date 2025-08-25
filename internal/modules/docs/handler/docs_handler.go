package handler

import (
	"docs-notify/internal/models"
	"docs-notify/internal/modules/docs/service"
	"docs-notify/internal/utils/exceptions"
	"errors"
	"net/http"

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
	// Parse form fields
	var doc models.Doc
	doc.UserId = c.Get("user_id").(uint) // from middleware
c.FormValue()
	// Bind JSON fields first (DocName, DocNo, etc.)
	if err := c.Bind(&doc); err != nil {
		return exceptions.NewResponseError(exceptions.ErrBadRequest, err)
	}
	if err := c.Validate(&doc); err != nil {
		return exceptions.NewResponseError(exceptions.ErrBadRequest, err)
	}

	// Expecting a single file with key "file"
	file, err := c.FormFile("file")
	if err != nil {
		return exceptions.NewResponseError(exceptions.ErrBadRequest, errors.New("Dokument faýlyny ýükläň"))
	}

	// Save in DB
	if err := h.docsService.CreateDoc(&doc, file); err != nil {
		return exceptions.NewResponseError(exceptions.ErrInternalServerError, errors.New("Dokument faýlyny ýükläň"))
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"message": "Document created successfully",
		"doc":     doc,
	})
}
