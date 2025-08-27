package dto

import (
	"time"
)

type DocResponse struct {
	ID           uint      `json:"id"`
	UserId       uint      `json:"user_id"`
	Username     string    `json:"username"`
	CategoryID   uint      `json:"category_id"`
	CategoryName string    `json:"category_name"`
	DocName      string    `json:"doc_name"`
	DocNo        string    `json:"doc_no"`
	EndDate      time.Time `json:"end_date"`
	NotifyDate   time.Time `json:"notify_date"`
	Status       string    `json:"status"`
	Permission   *uint     `json:"permission"` // from docs.permission
	UserPerm     *uint     `json:"user_perm"`  // from doc_users.permission
	File         string    `json:"file"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type DocsResponseDto struct {
	Docs  []DocResponse `json:"docs"`
	Total int64         `json:"total"`
}
