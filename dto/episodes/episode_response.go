package episodesdto

type EpisodeResponse struct {
	ID               int    `json:"id"`
	Title            string `json:"title" form:"title" gorm:"type: varchar(255)"`
	ThumbnailEpisode string `json:"thumbnailepisode" form:"thumbnailepisode" gorm:"type: varchar(255)"`
	LinkFilm         string `json:"linkepisode" form:"linkepisode" gorm:"type: varchar(255)"`
	FilmID           int    `json:"film_id" gorm:"type: int"`
	Description      string `json:"description" gorm:"type:text" form:"description"`
}
