package categoriesdto

type CreateCategoryRequest struct {
	Name      	   string          `json:"name" form:"name" gorm:"type: varchar(255)"`
}

type UpdateCategoryRequest struct {
	Name      	   string          `json:"name" form:"name" gorm:"type: varchar(255)"`
}
