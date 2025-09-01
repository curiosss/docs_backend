package handler

import (
	"docs-notify/internal/modules/notifications/repository"

	"docs-notify/internal/utils/exceptions"
	numutils "docs-notify/internal/utils/num_utils"
	"docs-notify/internal/utils/util"
	"net/http"

	"github.com/labstack/echo/v4"
)

type NotifsHandler struct {
	notifsRepo *repository.NotifsRepository
}

func NewNotifsHandler(repo *repository.NotifsRepository) *NotifsHandler {
	return &NotifsHandler{notifsRepo: repo}
}

func (h *NotifsHandler) Delete(c echo.Context) error {
	categId, err := numutils.GetUintParam(c, "category_id")
	if err != nil {
		return exceptions.NewResponseError(exceptions.ErrBadRequest, err)
	}

	err = h.notifsRepo.Delete(categId)
	if err != nil {
		return exceptions.NewResponseError(exceptions.ErrBadRequest, err)
	}

	return c.NoContent(http.StatusOK)
}

func (h *NotifsHandler) GetUserNotifications(c echo.Context) error {

	categs, err := h.notifsRepo.GetAll()
	if err != nil {
		return exceptions.NewResponseError(exceptions.ErrInternalServerError, err)
	}
	return c.JSON(http.StatusOK, util.WrapResponse(categs))
}

func (h *NotifsHandler) GetAdminNotifications(c echo.Context) error {

	categs, err := h.notifsRepo.GetAll()
	if err != nil {
		return exceptions.NewResponseError(exceptions.ErrInternalServerError, err)
	}
	return c.JSON(http.StatusOK, util.WrapResponse(categs))
}
