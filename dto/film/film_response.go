package filmdto

type FilmResponse struct {
	ID            int    `json:"id"`
	Title         string `json:"title" form:"title" gorm:"type: varchar(255)"`
	ThumbnailFilm string `json:"thumbnailfilm" form:"thumbnailfilm" gorm:"type: varchar(255)"`
	Year          string `json:"year" form:"year" gorm:"type: varchar(255)"`
	CategoryID    int    `json:"category_id" gorm:"type: int"`
	Description   string `json:"description" gorm:"type:text" form:"description"`
}
