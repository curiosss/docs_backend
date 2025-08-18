package handler

import (
	"docs-notify/internal/modules/users/dto"
	"docs-notify/internal/modules/users/service"
	"docs-notify/internal/utils/exceptions"
	numutils "docs-notify/internal/utils/num_utils"
	"docs-notify/internal/utils/util"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) Login(c echo.Context) error {

	var userLoginDto dto.UserLoginDto
	if err := c.Bind(&userLoginDto); err != nil {
		return exceptions.NewResponseError(exceptions.ErrBadRequest, err)
	}

	if err := c.Validate(&userLoginDto); err != nil {
		return exceptions.NewResponseError(exceptions.ErrBadRequest, err)
	}

	user, err := h.userService.Login(userLoginDto)
	if err != nil {
		return exceptions.NewResponseError(exceptions.ErrBadRequest, err)
	}

	return c.JSON(http.StatusOK, util.WrapResponse(user))
}

func (h *UserHandler) ChangeUsername(c echo.Context) error {

	userId := c.Get("user_id").(uint)
	var userLoginDto dto.UserLoginDto
	if err := c.Bind(&userLoginDto); err != nil {
		return exceptions.NewResponseError(exceptions.ErrBadRequest, err)
	}

	if err := c.Validate(&userLoginDto); err != nil {
		return exceptions.NewResponseError(exceptions.ErrBadRequest, err)
	}

	user, err := h.userService.ChangeUsername(&userLoginDto, userId)
	if err != nil {
		return exceptions.NewResponseError(exceptions.ErrBadRequest, err)
	}

	return c.JSON(http.StatusOK, util.WrapResponse(user))
}

func (h *UserHandler) ChangePassword(c echo.Context) error {
	userId := c.Get("user_id").(uint)
	var userPwdUpdateDto dto.UserPwdUpdateDto
	if err := c.Bind(&userPwdUpdateDto); err != nil {
		return exceptions.NewResponseError(exceptions.ErrBadRequest, err)
	}

	if err := c.Validate(&userPwdUpdateDto); err != nil {
		return exceptions.NewResponseError(exceptions.ErrBadRequest, err)
	}

	user, err := h.userService.ChangePassword(&userPwdUpdateDto, userId)
	if err != nil {
		return exceptions.NewResponseError(exceptions.ErrBadRequest, err)
	}

	return c.JSON(http.StatusOK, util.WrapResponse(user))
}

func (h *UserHandler) CreateUser(c echo.Context) error {
	var userCreateDto dto.UserCreateDto
	if err := c.Bind(&userCreateDto); err != nil {
		return exceptions.NewResponseError(exceptions.ErrBadRequest, err)
	}

	if err := c.Validate(&userCreateDto); err != nil {
		return exceptions.NewResponseError(exceptions.ErrBadRequest, err)
	}
	// fmt.Println("Creating user with data:", userCreateDto)
	user, err := h.userService.CreateUser(userCreateDto)
	if err != nil {
		return exceptions.NewResponseError(exceptions.ErrBadRequest, err)
	}

	return c.JSON(http.StatusCreated, util.WrapResponse(user))
}

func (h *UserHandler) UpdateUser(c echo.Context) error {
	var userUpdateDto dto.UserUpdateDto
	if err := c.Bind(&userUpdateDto); err != nil {
		return exceptions.NewResponseError(exceptions.ErrBadRequest, err)
	}

	if err := c.Validate(&userUpdateDto); err != nil {
		return exceptions.NewResponseError(exceptions.ErrBadRequest, err)
	}

	user, err := h.userService.UpdateUser(userUpdateDto)
	if err != nil {
		return exceptions.NewResponseError(exceptions.ErrBadRequest, err)
	}

	return c.JSON(http.StatusCreated, util.WrapResponse(user))
}

func (h *UserHandler) DeleteUser(c echo.Context) error {
	userId, err := numutils.GetUintParam(c, "user_id")
	if err != nil {
		return exceptions.NewResponseError(exceptions.ErrBadRequest, err)
	}

	err = h.userService.DeleteUser(userId)
	if err != nil {
		return exceptions.NewResponseError(exceptions.ErrBadRequest, err)
	}

	return c.NoContent(http.StatusOK)
}

func (h *UserHandler) GetAll(c echo.Context) error {

	users, err := h.userService.GetUsers()
	if err != nil {
		return exceptions.NewResponseError(exceptions.ErrInternalServerError, err)
	}
	return c.JSON(http.StatusCreated, util.WrapResponse(users))
}
