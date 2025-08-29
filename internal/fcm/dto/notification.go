package dto

type SendNotificationDto struct {
	Token string            `json:"token" validate:"required"`
	Title string            `json:"title" validate:"required"`
	Body  string            `json:"body" validate:"required"`
	Data  map[string]string `json:"data"`
}
