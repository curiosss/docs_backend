package dto

import "docs-notify/internal/models"

type DocsResponseDto struct{
	Data [] models.Doc
	Total int64
}