package repositories

import (
	"dumbflix/models"

	"gorm.io/gorm"
)

type FilmRepository interface {
	FindFilm() ([]models.Film, error)
	GetFilm(ID int) (models.Film, error)
	CreateFilm(film models.Film) (models.Film, error)
	UpdateFilm(film models.Film) (models.Film, error)
	DeleteFilm(film models.Film) (models.Film, error)
}

type repositoryForFilm struct {
	db *gorm.DB
}

func RepositoryFilm(db *gorm.DB) *repositoryForFilm {
	return &repositoryForFilm{db}
}

func (r *repositoryForFilm) FindFilm() ([]models.Film, error) {
	var films []models.Film
	err := r.db.Preload("Category").Find(&films).Error

	return films, err
}

func (r *repositoryForFilm) GetFilm(ID int) (models.Film, error) {
	var film models.Film
	err := r.db.First(&film, ID).Error

	return film, err
}

func (r *repositoryForFilm) CreateFilm(film models.Film) (models.Film, error) {
	err := r.db.Create(&film).Error

	return film, err
}

func (r *repositoryForFilm) UpdateFilm(film models.Film) (models.Film, error) {
	err := r.db.Save(&film).Error

	return film, err
}

func (r *repositoryForFilm) DeleteFilm(film models.Film) (models.Film, error) {
	err := r.db.Delete(&film).Error // Using Delete method

	return film, err
}
