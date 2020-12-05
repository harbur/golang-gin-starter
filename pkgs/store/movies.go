package store

import (
	"errors"

	"github.com/harbur/golang-gin-starter/pkgs/models"
	// sqlite
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// ListMovies lists movies
func (mgr *moviesDB) ListMovies() models.ListMovies {
	var movies []*models.Movie
	mgr.db.Find(&movies)
	return models.ListMovies{Movies: movies}
}

// CreateMovie creates a movie
func (mgr *moviesDB) CreateMovie(movie models.Movie) (models.Movie, error) {
	// make sure payload does not contain id
	if movie.ID != 0 {
		return movie, errors.New("invalid id")
	}
	err := mgr.db.Create(&movie).Error
	return movie, err
}

// GetMovie gets a movie
func (mgr *moviesDB) GetMovie(id uint) (models.GetMovie, error) {
	var movie models.Movie
	err := mgr.db.Where("ID = ?", id).Find(&movie).Error
	return models.GetMovie{Movie: movie}, err
}

// UpdateMovie updates a movie
func (mgr *moviesDB) UpdateMovie(id uint, movie models.Movie) (models.Movie, error) {
	// make sure payload contains correct id
	if id != movie.ID {
		return movie, errors.New("invalid id")
	}
	// make sure movie exists
	_, err := mgr.GetMovie(id)
	if err != nil {
		return movie, err
	}
	err = mgr.db.Save(&movie).Error
	return movie, err
}

// DeleteMovie deletes a movies
func (mgr *moviesDB) DeleteMovie(id uint) {
	var movie models.Movie
	mgr.db.Where("ID = ?", id).Delete(movie)
}
