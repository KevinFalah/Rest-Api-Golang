package routes

import (
	"dumbflix/handlers"
	"dumbflix/pkg/mysql"
	"dumbflix/repositories"

	"github.com/gorilla/mux"
)

func FilmRoutes(r *mux.Router) {
	FilmRepository := repositories.RepositoryFilm(mysql.DB)
	h := handlers.HandlerFilm(FilmRepository)

	r.HandleFunc("/films", h.FindFilms).Methods("GET")
	r.HandleFunc("/film/{id}", h.GetFilm).Methods("GET")
	r.HandleFunc("/film", h.CreateFilm).Methods("POST")
	r.HandleFunc("/film/{id}", h.DeleteFilm).Methods("DELETE")
}
