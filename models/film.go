package models

import "time"

type Film struct {
	ID            int                `json:"id" gorm:"primary_key:auto_increment"`
	Title         string             `json:"title" form:"title" gorm:"type: varchar(255)"`
	ThumbnailFilm string             `json:"thumbnailfilm" form:"thumbnailfilm" gorm:"type: varchar(255)"`
	Year          string             `json:"year" form:"year" gorm:"type: varchar(255)"`
	CategoryID    int                `json:"category_id" form:"category_id"`
	Category      CategoriesResponse `json:"category"`
	Description   string             `json:"description" gorm:"type:text" form:"description"`
	// Category   []Category           `json:"category" gorm:"many2many:film_categories"`
	// CategoryID []int                `json:"-" form:"category_id" gorm:"-"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type FilmResponse struct {
	ID            int                `json:"id"`
	Title         string             `json:"title"`
	ThumbnailFilm string             `json:"thumbnailfilm"`
	Year          string             `json:"year"`
	CategoryID    int                `json:"-"`
	Category      CategoriesResponse `json:"category"`
	Description   string             `json:"description"`
	// Category   []Category           `json:"category" gorm:"many2many:film_categories"`
	// CategoryID []int                `json:"-" form:"category_id" gorm:"-"`
}

type FilmEpisodeResponse struct {
	ID            int                `json:"id"`
	Title         string             `json:"title"`
	ThumbnailFilm string             `json:"thumbnailfilm"`
	Year          string             `json:"year"`
	CategoryID    int                `json:"-"`
	Category      CategoriesResponse `json:"category"`
	// Description   string             `json:"description"`
	// Category   []Category           `json:"category" gorm:"many2many:film_categories"`
	// CategoryID []int                `json:"-" form:"category_id" gorm:"-"`
}

type FilmCategoryResponse struct {
	ID            int    `json:"id"`
	Title         string `json:"title"`
	ThumbnailFilm string `json:"thumbnailfilm"`
	Year          string `json:"year"`
	CategoryID    int    `json:"-"`
	Description   string `json:"description"`
}

func (FilmEpisodeResponse) TableName() string {
	return "films"
}

func (FilmResponse) TableName() string {
	return "films"
}

func (FilmCategoryResponse) TableName() string {
	return "films"
}
