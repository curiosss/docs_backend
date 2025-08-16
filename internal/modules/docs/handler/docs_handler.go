package handler

import (
	"docs-notify/internal/models"
	"docs-notify/internal/modules/docs/service"
	"docs-notify/internal/utils/exceptions"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"

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
	doc.UserId = c.Get("user_id").(uint) // Assuming user_id is set in AuthMiddleware
	if err := c.Bind(&doc); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid request"})
	}
	if err := c.Validate(&doc); err != nil {
		return exceptions.NewResponseError(exceptions.ErrBadRequest, err)
	}
	fmt.Println("Received doc:", doc)
	// Handle file uploads
	form, err := c.MultipartForm()
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "failed to parse form"})
	}

	files := form.File["files"] // frontend should send with key "files"
	var savedFiles []models.File

	for _, file := range files {
		src, err := file.Open()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": "failed to open file"})
		}
		defer src.Close()

		// Save to local disk (you can swap with S3/Google Cloud later)
		uploadDir := "uploads"
		if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
			os.Mkdir(uploadDir, os.ModePerm)
		}
		filename := fmt.Sprintf("%d_%s", time.Now().UnixNano(), file.Filename)
		filepath := filepath.Join(uploadDir, filename)

		dst, err := os.Create(filepath)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": "failed to save file"})
		}
		defer dst.Close()

		if _, err = dst.ReadFrom(src); err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": "failed to write file"})
		}

		savedFiles = append(savedFiles, models.File{
			Filename: file.Filename,
			Filepath: filepath,
			URL:      "/static/" + filename, // Example: expose via /static/
		})
	}

	// Save doc + files in service
	if err := h.docsService.CreateDocWithFiles(&doc, savedFiles); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"message": "Document created successfully",
		"doc":     doc,
		"files":   savedFiles,
	})
}
