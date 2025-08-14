package dto

type UserResponseDto struct {
	ID        uint   `json:"id"`
	Username  string `json:"username"`
	Role      string `json:"role"`
	Note      string `json:"note"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
