package dto

// DTO для остальных сущностей будут здесь
// Например, для Doc:

type CreateDocRequest struct {
	Title      string `json:"title" validate:"required"`
	Content    string `json:"content"`
	CategoryID uint   `json:"category_id"`
	NotifyDate string `json:"notify_date"` // format: "2006-01-02T15:04:05Z07:00"
}

type DocResponse struct {
	ID         uint   `json:"id"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	AuthorID   uint   `json:"author_id"`
	CategoryID uint   `json:"category_id"`
	NotifyDate string `json:"notify_date"`
}
