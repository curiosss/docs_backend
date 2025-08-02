package handlers

import (
	"net/http"

	"docs-notify/internal/dto"
	"docs-notify/internal/services"
	"docs-notify/internal/utils"
	"docs-notify/internal/validators"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	service services.UserService
}

func NewUserHandler(s services.UserService) *UserHandler {
	return &UserHandler{service: s}
}

// Register godoc
// @Summary      Register a new user
// @Description  Creates a new user account
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        user body dto.RegisterUserRequest true "User registration data"
// @Success      201  {object}  utils.SuccessResponse{data=dto.UserResponse} "User created successfully"
// @Failure      400  {object}  utils.ErrorResponse "Invalid input"
// @Failure      500  {object}  utils.ErrorResponse "Internal server error"
// @Router       /users/register [post]
func (h *UserHandler) Register(c echo.Context) error {
	var req dto.RegisterUserRequest
	if err := c.Bind(&req); err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request body")
	}

	if err := validators.ValidateRegisterUser(&req); err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	user, err := h.service.Register(&req)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return utils.SuccessResponse(c, http.StatusCreated, "User registered successfully", user)
}

// Login godoc
// @Summary      Login a user
// @Description  Authenticates a user and returns a JWT token
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        credentials body dto.LoginUserRequest true "User login credentials"
// @Success      200  {object}  utils.SuccessResponse{data=dto.LoginResponse} "Login successful"
// @Failure      400  {object}  utils.ErrorResponse "Invalid input"
// @Failure      401  {object}  utils.ErrorResponse "Invalid credentials"
// @Router       /users/login [post]
func (h *UserHandler) Login(c echo.Context) error {
	var req dto.LoginUserRequest
	if err := c.Bind(&req); err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request body")
	}

	if err := validators.ValidateLoginUser(&req); err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	response, err := h.service.Login(&req)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusUnauthorized, err.Error())
	}

	return utils.SuccessResponse(c, http.StatusOK, "Login successful", response)
}
