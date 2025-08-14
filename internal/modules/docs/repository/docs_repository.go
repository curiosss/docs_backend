package repository

import (
	"gorm.io/gorm"
)

type DocsRepository struct {
	db *gorm.DB
}

func NewDocsRepository(db *gorm.DB) *DocsRepository {
	return &DocsRepository{db: db}
}
