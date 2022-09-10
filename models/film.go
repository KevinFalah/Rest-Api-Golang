package models

import "time"

type Film struct {
	ID            int              `json:"id" gorm:"primary_key:auto_increment"`
	Title         string           `json:"title" form:"title" gorm:"type: varchar(255)"`
	ThumbnailFilm string           `json:"thumbnailfilm" form:"thumbnailfilm" gorm:"type: varchar(255)"`
	Image         string           `json:"image" form:"image" gorm:"type: varchar(255)"`
	Description   string           `json:"description" gorm:"type:text" form:"desc"`
	Year          int              `json:"year" gorm:"type: int"`
	Category      CategoryResponse `json:"category"`
	CategoryID    int              `json:"category_id" form:"category_id"`
	CreatedAt     time.Time        `json:"-"`
	UpdatedAt     time.Time        `json:"-"`
	// User          UserIdResponse `json:"user"`
	// UserID        int              `json:"user_id" form:"user_id"`
}

// type FilmResponse struct {
// 	ID            int    `json:"id"`
// 	Title         string `json:"title"`
// 	ThumbnailFilm string `json:"thumbnailfilm"`
// 	Image         string `json:"image"`
// 	Year          int    `json:"year"`
// }

// func (FilmResponse) TableName() string {
// 	return "films"
// }

// type ProductResponse struct {
//   ID         int                  `json:"id"`
//   Name       string               `json:"name"`
//   Desc       string               `json:"desc"`
//   Price      int                  `json:"price"`
//   Image      string               `json:"image"`
//   Qty        int                  `json:"qty"`
//   UserID     int                  `json:"-"`
//   User       UsersProfileResponse `json:"user"`
//   Category   []Category           `json:"category" gorm:"many2many:product_categories"`
//   CategoryID []int                `json:"category_id" form:"category_id" gorm:"-"`
// }

// type ProductUserResponse struct {
//   ID     int    `json:"id"`
//   Name   string `json:"name"`
//   Desc   string `json:"desc"`
//   Price  int    `json:"price"`
//   Image  string `json:"image"`
//   Qty    int    `json:"qty"`
//   UserID int    `json:"-"`
// }

// func (ProductResponse) TableName() string {
//   return "products"
// }

// func (ProductUserResponse) TableName() string {
//   return "products"
// }
