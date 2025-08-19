package handler

import (
	"docs-notify/internal/models"
	"docs-notify/internal/modules/categories/dto"
	"docs-notify/internal/modules/categories/repository"
	"docs-notify/internal/utils/exceptions"
	numutils "docs-notify/internal/utils/num_utils"
	"docs-notify/internal/utils/util"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CategoriesHandler struct {
	categRepo *repository.CategoriesRepository
}

func NewCategHandler(repo *repository.CategoriesRepository) *CategoriesHandler {
	return &CategoriesHandler{categRepo: repo}
}

func (h *CategoriesHandler) Create(c echo.Context) error {
	var categoryCreateDto dto.CategoryCreateDto
	if err := c.Bind(&categoryCreateDto); err != nil {
		return exceptions.NewResponseError(exceptions.ErrBadRequest, err)
	}

	if err := c.Validate(&categoryCreateDto); err != nil {
		return exceptions.NewResponseError(exceptions.ErrBadRequest, err)
	}
	categ, err := h.categRepo.Create(&categoryCreateDto)
	if err != nil {
		return exceptions.NewResponseError(exceptions.ErrBadRequest, err)
	}

	return c.JSON(http.StatusCreated, util.WrapResponse(categ))
}

func (h *CategoriesHandler) Update(c echo.Context) error {
	var category models.Category
	if err := c.Bind(&category); err != nil {
		return exceptions.NewResponseError(exceptions.ErrBadRequest, err)
	}

	if err := c.Validate(&category); err != nil {
		return exceptions.NewResponseError(exceptions.ErrBadRequest, err)
	}

	categ, err := h.categRepo.Update(&category)
	if err != nil {
		return exceptions.NewResponseError(exceptions.ErrBadRequest, err)
	}

	return c.JSON(http.StatusOK, util.WrapResponse(categ))
}

func (h *CategoriesHandler) Delete(c echo.Context) error {
	categId, err := numutils.GetUintParam(c, "category_id")
	if err != nil {
		return exceptions.NewResponseError(exceptions.ErrBadRequest, err)
	}

	err = h.categRepo.Delete(categId)
	if err != nil {
		return exceptions.NewResponseError(exceptions.ErrBadRequest, err)
	}

	return c.NoContent(http.StatusOK)
}

func (h *CategoriesHandler) GetAll(c echo.Context) error {

	categs, err := h.categRepo.GetAll()
	if err != nil {
		return exceptions.NewResponseError(exceptions.ErrInternalServerError, err)
	}
	return c.JSON(http.StatusOK, util.WrapResponse(categs))
}
