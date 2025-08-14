package service

import (
	"docs-notify/internal/config"
	"docs-notify/internal/modules/docs/repository"
)

type DocsService struct {
	userRepository *repository.DocsRepository
	config         *config.Config
}

func NewDocsService(userRepository *repository.DocsRepository, cfg *config.Config) *DocsService {
	return &DocsService{userRepository: userRepository, config: cfg}
}
