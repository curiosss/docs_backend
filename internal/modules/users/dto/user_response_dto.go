package dto

type UserResponseDto struct {
	ID           uint   `json:"id"`
	Username     string `json:"username"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
