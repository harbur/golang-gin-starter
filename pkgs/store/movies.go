package store

import (
	"errors"

	"github.com/harbur/golang-gin-starter/pkgs/models"
	// sqlite
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// ListMovies lists movies
func ListMovies() models.Movies {
	var movies models.Movies
	db.Find(&movies)
	return movies
}

// CreateMovie creates a movie
func CreateMovie(movie models.Movie) (models.Movie, error) {
	// make sure payload does not contain id
	if movie.ID != 0 {
		return movie, errors.New("invalid id")
	}
	err := db.Create(&movie).Error
	return movie, err
}

// GetMovie gets a movie
func GetMovie(id int64) (models.Movie, error) {
	var movie models.Movie
	err := db.Where("ID = ?", id).Find(&movie).Error
	return movie, err
}

// UpdateMovie updates a movie
func UpdateMovie(id int64, movie models.Movie) (models.Movie, error) {
	// make sure payload contains correct id
	if id != movie.ID {
		return movie, errors.New("invalid id")
	}
	// make sure movie exists
	_, err := GetMovie(id)
	if err != nil {
		return movie, err
	}
	err = db.Save(&movie).Error
	return movie, err
}

// DeleteMovie deletes a movies
func DeleteMovie(id int64) {
	var movie models.Movie
	db.Where("ID = ?", id).Delete(movie)
}
