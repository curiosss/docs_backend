package validator

import (
	"docs-notify/internal/dto"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func ValidateRegisterUser(req *dto.RegisterUserRequest) error {
	return validate.Struct(req)
}

func ValidateLoginUser(req *dto.LoginUserRequest) error {
	return validate.Struct(req)
}
