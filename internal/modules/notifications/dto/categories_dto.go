package dto

type CategoryCreateDto struct {
	Name     string `json:"name" gorm:"unique;not null;size:100" validate:"required"`
	ParentID *uint  `json:"parent_id" gorm:"default:null"` // For subcategories

}
