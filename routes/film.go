package routes

import (
	"dumbflix/handlers"
	"dumbflix/pkg/mysql"
	"dumbflix/repositories"

	"github.com/gorilla/mux"
)

func filmRoutes(r *mux.Router) {
	filmRepository := repositories.RepositoryFilm(mysql.DB)
	h := handlers.HandlerFilm(filmRepository)

	r.HandleFunc("/films", h.FindFilm).Methods("GET")
	r.HandleFunc("/film/{id}", h.GetFilm).Methods("GET")
	r.HandleFunc("/film", h.CreateFilm).Methods("POST")
	r.HandleFunc("/film/{id}", h.UpdateFilm).Methods("PATCH")
	r.HandleFunc("/film/{id}", h.DeleteFilm).Methods("DELETE")
}
