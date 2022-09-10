package filmsdto

type CreateFilmRequest struct {
	Title         string `json:"title" form:"title" gorm:"type: varchar(255)"`
	ThumbnailFilm string `json:"thumbnailfilm" form:"thumbnailfilm" gorm:"type: varchar(255)"`
	Description   string `json:"description" gorm:"type:text" form:"desc"`
	Year          int    `json:"year" form:"year" gorm:"type: int"`
	CategoryID   int    `json:"category"`
}

type UpdateFilmRequest struct {
	Title         string `json:"title" form:"name" gorm:"type: varchar(255)"`
	ThumbnailFilm string `json:"thumbnailFilm" form:"image" gorm:"type: varchar(255)"`
	Description   string `json:"description" gorm:"type:text" form:"desc"`
	Year          int    `json:"year" form:"price" gorm:"type: int"`
	CategoryID    int    `json:"category"`
}
