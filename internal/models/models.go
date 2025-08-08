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
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Category struct {
	gorm.Model
	Name string `gorm:"unique;not null;size:100"`
	Docs []Doc
}

type Doc struct {
	gorm.Model
	Title      string `gorm:"not null;size:255"`
	Content    string
	CategoryID uint
	AuthorID   uint `gorm:"not null"`
	NotifyDate time.Time
	NotifSent  bool   `gorm:"default:false"`
	Users      []User `gorm:"many2many:doc_users;"`
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
