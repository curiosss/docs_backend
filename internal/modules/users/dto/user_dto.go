package dto

type UserLoginDto struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserCreateDto struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Role     string `json:"role" validate:"required,oneof=admin operator user"`
	Note     string `json:"note" validate:"omitempty"`
}

type UserUpdateDto struct {
	Id       uint   `json:"id" validate:"required"`
	Username string `json:"username" validate:"omitempty"`
	Password string `json:"password" validate:"omitempty"`
	Role     string `json:"role" validate:"omitempty,oneof=admin operator user"`
	Note     string `json:"note" validate:"omitempty"`
}
