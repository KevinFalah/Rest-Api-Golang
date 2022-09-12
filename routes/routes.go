package routes

import (
	"github.com/gorilla/mux"
)

func RouteInit(r *mux.Router) {
	CategoryRoutes(r)
	FilmRoutes(r)
	EpisodeRoutes(r)
	TransactionRoutes(r)
	UserRoutes(r)
	AuthRoutes(r)
}
