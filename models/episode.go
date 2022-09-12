package models

import "time"

type Episode struct {
	ID               int                 `json:"id"`
	Title            string              `json:"title" form:"title" gorm:"type: varchar(255)"`
	ThumbnailEpisode string              `json:"thumbnailepisode" form:"thumbnailepisode" gorm:"type: varchar(255)"`
	LinkFilm         string              `json:"linkfilm" form:"linkfilm" gorm:"type: varchar(255)"`
	FilmID           int                 `json:"film_id" form:"film_id"`
	Film             FilmEpisodeResponse `json:"film"`
	Description      string              `json:"description" gorm:"type:text" form:"description"`
	CreatedAt        time.Time           `json:"-"`
	UpdatedAt        time.Time           `json:"-"`
}

// type EpisodeE struct {
// 	ID               int                 `json:"id"`
// 	Title            string              `json:"title" form:"title" gorm:"type: varchar(255)"`
// 	ThumbnailEpisode string              `json:"thumbnailepisode" form:"thumbnailepisode" gorm:"type: varchar(255)"`
// 	LinkFilm         string              `json:"linkfilm" form:"linkfilm" gorm:"type: varchar(255)"`
// 	FilmID           int                 `json:"film_id" form:"film_id"`
// 	Film             FilmEpisodeResponse `json:"film"`
// 	Description      string              `json:"description" gorm:"type:text" form:"description"`
// 	CreatedAt        time.Time           `json:"-"`
// 	UpdatedAt        time.Time           `json:"-"`
// }

type EpisodeResponse struct {
	ID               int                 `json:"id"`
	Title            string              `json:"title"`
	ThumbnailEpisode string              `json:"image"`
	LinkFilm         string              `json:"linkfilm"`
	FilmID           int                 `json:"-"`
	Film             FilmEpisodeResponse `json:"film"`
	Description      string              `json:"description"`
}

// type EpisodeResponseFilm struct {
// 	ID               int                 `json:"id"`
// 	Title            string              `json:"title"`
// 	ThumbnailEpisode string              `json:"image"`
// 	LinkFilm         string              `json:"linkfilm"`
// 	FilmID           int                 `json:"-"`
// 	Film             FilmEpisodeResponse `json:"film"`
// 	Description      string              `json:"description"`
// }

func (EpisodeResponse) TableName() string {
	return "episodes"
}

// func (EpisodeResponseFilm) TableName() string {
// 	return "episodes"
// }

// func (EpisodeE) TableName() string {
// 	return "episodes"
// }
