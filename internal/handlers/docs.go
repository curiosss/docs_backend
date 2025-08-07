package handlers

import (
	"net/http"
	"strconv"

	"docs-notify/internal/dto"
	"docs-notify/internal/services"
	"docs-notify/internal/utils"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type DocHandler struct {
	service services.DocService
}

func NewDocHandler(s services.DocService) *DocHandler {
	return &DocHandler{service: s}
}

// CreateDoc godoc
// @Summary      Create a new document
// @Description  Adds a new document to the system
// @Tags         docs
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        doc body dto.CreateDocRequest true "Document data"
// @Success      201  {object}  utils.SuccessResponse{data=dto.DocResponse} "Document created"
// @Failure      400  {object}  utils.ErrorResponse "Invalid input"
// @Failure      401  {object}  utils.ErrorResponse "Unauthorized"
// @Failure      500  {object}  utils.ErrorResponse "Internal server error"
// @Router       /docs [post]
func (h *DocHandler) CreateDoc(c echo.Context) error {
	var req dto.CreateDocRequest
	if err := c.Bind(&req); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request body")
	}

	// if err := validators.ValidateCreateDoc(&req); err != nil {
	// 	return utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
	// }

	user := c.Get("user").(jwt.MapClaims)
	authorID := uint(user["user_id"].(float64))

	doc, err := h.service.CreateDoc(&req, authorID)
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, err.Error())
	}

	return utils.SendSuccess(c, http.StatusCreated, "Document created successfully", doc)
}

func (h *DocHandler) GetDoc(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid document ID")
	}

	doc, err := h.service.GetDoc(uint(id))
	if err != nil {
		// Здесь можно добавить проверку на gorm.ErrRecordNotFound для 404
		return utils.SendError(c, http.StatusInternalServerError, err.Error())
	}

	return utils.SendSuccess(c, http.StatusOK, "Document retrieved successfully", doc)
}
