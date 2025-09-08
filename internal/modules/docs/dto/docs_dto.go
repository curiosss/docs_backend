package dto

type DocCreateDto struct {
	UserId        uint
	CategoryID    uint    `form:"category_id" gorm:"not null" validate:"required"`
	SubCategoryId *uint   `form:"sub_category_id"`
	DocName       string  `form:"doc_name" gorm:"not null;size:255" validate:"required"`
	DocNo         string  `form:"doc_no" gorm:"not null;size:100" validate:"required"`
	EndDate       string  `form:"end_date" validate:"required"`
	NotifyDate    string  `form:"notify_date" validate:"required"`
	Status        string  `form:"status" gorm:"not null;default:'pending'" validate:"required"`
	Permission    *uint   `form:"permission"`  // 0: private, 1: public
	Permissions   *string `form:"permissions"` // List of user IDs with access

}

type DocUpdateDto struct {
	Id            uint    `form:"id" gorm:"primaryKey;autoIncrement" validate:"required"`
	CategoryID    uint    `form:"category_id" gorm:"not null" validate:"required"`
	SubCategoryId *uint   `form:"sub_category_id"`
	DocName       string  `form:"doc_name" gorm:"not null;size:255" validate:"required"`
	DocNo         string  `form:"doc_no" gorm:"not null;size:100" validate:"required"`
	EndDate       string  `form:"end_date" validate:"required"`
	NotifyDate    string  `form:"notify_date" validate:"required"`
	Status        string  `form:"status" validate:"required"`
	Permission    *uint   `form:"permission"`  // 0: private, 1: public
	Permissions   *string `form:"permissions"` // List of user IDs with access
}

type GetDocsDto struct {
	Page          int `query:"page" default:"1" validate:"min=1"`
	Limit         int `query:"limit" default:"20" validate:"min=1,max=100"`
	UserId        uint
	Categories    []uint  `query:"categories"`
	CategoryID    *uint   `query:"category_id"`
	SubCategoryId *uint   `query:"sub_category_id"`
	Status        *string `query:"status"`
	CreatedUserId *uint   `query:"created_user_id"`
	CreatedFrom   *string `query:"created_from"`
	CreatedTo     *string `query:"created_to"`
	PreparedFrom  *string `query:"prepared_from"`
	PreparedTo    *string `query:"prepared_to"`
	SearchText    *string `query:"search_text"`
}

type GetDocStatsDto struct {
	DateType string  `query:"date_type"`
	DateFrom *string `query:"date_from"`
	DateTo   *string `query:"date_to"`
	UserIds  []uint  `query:"user_ids"`
}

// type DocUserPm struct {
// 	UserId uint
// 	DocId uin
// }
