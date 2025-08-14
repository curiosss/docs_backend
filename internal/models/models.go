package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID          uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Username    string    `json:"username" gorm:"unique;not null;size:50"`
	Password    string    `json:"password" gorm:"not null"`
	Role        string    `json:"role" gorm:"not null"`
	FcmToken    string    `json:"fcm_token"`
	AccessToken string    `json:"access_token"`
	Note        string    `json:"note"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Doc struct {
	ID          uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	UserId      uint      `json:"user_id" gorm:"not null"`
	CategoryID  uint      `json:"category_id" gorm:"not null"`
	DocName     string    `json:"doc_name" gorm:"not null;size:255"`
	DocNo       string    `json:"doc_no" gorm:"not null;size:100"`
	EndDate     time.Time `json:"end_date" `
	NotifyDate  time.Time `json:"notify_date"`
	NotifSent   bool      `gorm:"default:false"`
	Status      string    `json:"status" gorm:"not null;default:'active'"`
	Perminssion uint      `json:"perminssion" gorm:"not null;default:0"` // 0: private, 1: public
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Category struct {
	gorm.Model
	Name string `gorm:"unique;not null;size:100"`
	Docs []Doc
}

type DocUser struct {
	DocID    uint      `gorm:"primaryKey"`
	UserID   uint      `gorm:"primaryKey"`
	SharedAt time.Time `gorm:"autoCreateTime"`
}

type Action struct {
	gorm.Model
	UserID     uint
	DocID      uint
	ActionType string `gorm:"size:50;not null"` // e.g., 'created', 'updated', 'viewed', 'shared'
}

type File struct {
	gorm.Model
	DocID    uint
	Filename string `gorm:"size:255;not null"`
	Filepath string `gorm:"size:255;not null"`
	URL      string `gorm:"size:255;not null"`
}
