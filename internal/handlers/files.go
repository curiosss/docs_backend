package handlers

// import (
// 	"fmt"
// 	"io"
// 	"net/http"
// 	"os"
// 	"path/filepath"
// 	"strconv"

// 	"docs-notify/internal/services"
// 	"docs-notify/internal/utils"

// 	"github.com/labstack/echo/v4"
// )

// type FileHandler struct {
// 	docService services.DocService // Используем DocService для работы с документами
// }

// func NewFileHandler(ds services.DocService) *FileHandler {
// 	return &FileHandler{docService: ds}
// }

// func (h *FileHandler) UploadFile(c echo.Context) error {
// 	// --- Получение файла из запроса ---
// 	file, err := c.FormFile("file")
// 	if err != nil {
// 		return utils.ErrorResponse(c, http.StatusBadRequest, "File is missing")
// 	}

// 	// --- Получение ID документа ---
// 	docIDStr := c.FormValue("doc_id")
// 	docID, err := strconv.Atoi(docIDStr)
// 	if err != nil {
// 		return utils.ErrorResponse(c, http.StatusBadRequest, "Invalid doc_id")
// 	}

// 	// --- Сохранение файла ---
// 	src, err := file.Open()
// 	if err != nil {
// 		return utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to open file")
// 	}
// 	defer src.Close()

// 	// Создаем директорию, если ее нет
// 	uploadDir := "./uploads"
// 	if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
// 		os.Mkdir(uploadDir, 0o755)
// 	}

// 	// Создаем путь для сохранения
// 	filePath := filepath.Join(uploadDir, file.Filename)
// 	dst, err := os.Create(filePath)
// 	if err != nil {
// 		return utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to create file on server")
// 	}
// 	defer dst.Close()

// 	// Копируем содержимое
// 	if _, err = io.Copy(dst, src); err != nil {
// 		return utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to save file")
// 	}

// 	// --- Сохранение метаданных в БД ---
// 	publicURL := fmt.Sprintf("/static/%s", file.Filename)

// 	// Здесь вы должны вызвать метод сервиса для сохранения информации о файле в БД
// 	// Например: fileInfo, err := h.docService.AddFileToDoc(uint(docID), file.Filename, filePath, publicURL)
// 	// Для примера, просто вернем URL

// 	fileInfo := map[string]string{
// 		"filename": file.Filename,
// 		"url":      publicURL,
// 	}

// 	return utils.SuccessResponse(c, http.StatusOK, "File uploaded successfully", fileInfo)
// }
