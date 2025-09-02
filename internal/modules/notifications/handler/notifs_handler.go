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

func (h *NotifsHandler) GetUserNotifications(c echo.Context) error {

	userId := c.Get("user_id").(uint)
	page, err := numutils.GetIntParam(c, "page")
	if err != nil {
		return exceptions.NewResponseError(exceptions.ErrBadRequest, err)
	}

	items, err := h.notifsRepo.GetAll(userId, page)
	if err != nil {
		return exceptions.NewResponseError(exceptions.ErrInternalServerError, err)
	}
	return c.JSON(http.StatusOK, util.WrapResponse(items))
}

func (h *NotifsHandler) GetAdminNotifications(c echo.Context) error {

	// categs, err := h.notifsRepo.GetAll()
	// if err != nil {
	// 	return exceptions.NewResponseError(exceptions.ErrInternalServerError, err)
	// }
	// return c.JSON(http.StatusOK, util.WrapResponse(categs))
	return nil
}
