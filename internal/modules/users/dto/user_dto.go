package dto

type UserLoginDto struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserUpdateDto struct {
	Username string `json:"username" validate:"omitempty"`
	Name     string `json:"name" validate:"omitempty"`
	Password string `json:"password" validate:"omitempty"`
	Email    string `json:"email" validate:"omitempty,email"`
}
