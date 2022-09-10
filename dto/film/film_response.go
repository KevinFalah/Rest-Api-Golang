package filmsdto

type FilmResponse struct {
	ID            int    `json:"id" gorm:"primary_key:auto_increment"`
	Title         string `json:"title" form:"title" gorm:"type: varchar(255)"`
	ThumbnailFilm string `json:"thumbnailfilm" form:"thumbnailfilm" gorm:"type: varchar(255)"`
	Description   string `json:"desc" gorm:"type:text" form:"desc"`
	Year          int    `json:"year" form:"price" gorm:"type: int"`
	CategoryID    int    `json:"category"`
}
