package episodesdto

type EpisodeRequest struct {
	Title            string `json:"title" form:"title" gorm:"type: varchar(255)"`
	ThumbnailEpisode string `json:"thumbnailepisode" form:"thumbnailepisode" gorm:"type: varchar(255)"`
	LinkFilm         string `json:"linkfilm" form:"linkfilm" gorm:"type: varchar(255)"`
	FilmID           int    `json:"film_id" gorm:"type: int"`
	Description      string `json:"description" gorm:"type:text" form:"description"`
}

type CreateEpisodeRequest struct {
	Title            string `json:"title" form:"title"`
	ThumbnailEpisode string `json:"thumbnailepisode" form:"thumbnailepisode"`
	LinkFilm         string `json:"linkfilm" form:"linkfilm"`
	FilmID           int    `json:"film_id"`
	Description      string `json:"description" gorm:"type:text" form:"description"`
}

type UpdateEpisodeRequest struct {
	Title            string `json:"title" form:"title"`
	ThumbnailEpisode string `json:"thumbnailepisode" form:"thumbnailepisode"`
	LinkFilm         string `json:"linkfilm" form:"linkfilm"`
	FilmID           int    `json:"film_id"`
	Description      string `json:"description" gorm:"type:text" form:"description"`
}
