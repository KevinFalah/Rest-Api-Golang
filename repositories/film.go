package repositories

import (
	"dumbflix/models"

	"gorm.io/gorm"
)

type FilmRepository interface {
	FindFilms() ([]models.Film, error)
	GetFilm(ID int) (models.Film, error)
	CreateFilm(Film models.Film) (models.Film, error)
	UpdateFilm(Film models.Film) (models.Film, error)
	DeleteFilm(Film models.Film) (models.Film, error)
}

func RepositoryFilm(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindFilms() ([]models.Film, error) {
	var films []models.Film
	err := r.db.Preload("Category").Find(&films).Error // add this code

	return films, err
}

func (r *repository) GetFilm(ID int) (models.Film, error) {
	var Film models.Film
	// not yet using category relation, cause this step doesnt Belong to Many
	err := r.db.Preload("Category").First(&Film, ID).Error // add this code

	return Film, err
}

func (r *repository) CreateFilm(Film models.Film) (models.Film, error) {
	err := r.db.Create(&Film).Error

	return Film, err
}

func (r *repository) UpdateFilm(Film models.Film) (models.Film, error) {
	err := r.db.Save(&Film).Error

	return Film, err
}

func (r *repository) DeleteFilm(Film models.Film) (models.Film, error) {
	err := r.db.Delete(&Film).Error

	return Film, err
}
