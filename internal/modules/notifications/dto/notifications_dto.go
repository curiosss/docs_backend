package dto

type NotificationAdminResponseDto struct {
	ID       uint   `json:"id"`
	DocID    uint   `json:"doc_id"`
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	UserRole string `json:"user_role"`
	IsSeen   bool   `json:"is_seen"`
	Title    string `json:"title"`
	Body     string `json:"body"`
}
