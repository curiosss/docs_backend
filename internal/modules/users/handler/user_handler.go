package handler

import (
	"docs-notify/internal/modules/users/dto"
	"docs-notify/internal/modules/users/service"
	"docs-notify/internal/utils/exceptions"
	"docs-notify/internal/utils/util"
	"fmt"
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
		fmt.Println(err)
		// return err
		return exceptions.NewResponseError(exceptions.ErrBadRequest, err)
	}

	return c.JSON(http.StatusOK, util.WrapResponse(user))
}

// func (h *AuthHandler) Register(c echo.Context) error {

// 	var registerDto dto.UserRegisterDto
// 	if err := c.Bind(&registerDto); err != nil {
// 		return exceptions.NewResponseError(exceptions.ErrBadRequest, err)
// 	}

// 	if err := c.Validate(&registerDto); err != nil {
// 		return exceptions.NewResponseError(exceptions.ErrBadRequest, err)
// 	}

// 	admin, err := h.userService.Register(registerDto)
// 	if err != nil {
// 		return exceptions.NewResponseError(exceptions.ErrBadRequest, err)
// 	}
// 	return c.JSON(http.StatusCreated, util.WrapResponse(admin))
// }

// func (h *AuthHandler) GetAll(c echo.Context) error {
// 	users, err := h.userService.GetAll()
// 	if err != nil {
// 		return exceptions.NewResponseError(exceptions.ErrBadRequest, err)
// 	}
// 	return c.JSON(http.StatusOK, util.WrapResponse(users))
// }
// func (h *AuthHandler) Update(c echo.Context) error {

// 	var updateDto dto.UserUpdateDto
// 	if err := c.Bind(&updateDto); err != nil {
// 		return exceptions.NewResponseError(exceptions.ErrBadRequest, err)
// 	}

// 	if err := c.Validate(&updateDto); err != nil {
// 		return exceptions.NewResponseError(exceptions.ErrBadRequest, err)
// 	}

// 	userId := c.Get("user_id").(uint)
// 	user, err := h.userService.Update(userId, updateDto)
// 	if err != nil {
// 		return exceptions.NewResponseError(exceptions.ErrBadRequest, err)
// 	}

// 	return c.JSON(http.StatusOK, util.WrapResponse(user))
// }
