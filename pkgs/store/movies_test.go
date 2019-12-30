package store

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/harbur/golang-starter/models"
	"github.com/harbur/golang-starter/pkgs/utils"
)

// TestCreateMovieOK creates a movie correctly
func TestCreateMovieOK(t *testing.T) {
	// Prepare
	Connect()
	movie := models.Movie{
		Name: utils.StrPtr("godfather"),
	}

	// Command
	movie, err := CreateMovie(movie)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, int64(1), movie.ID, "the movie ID should now be 1")
	var stored models.Movie
	db.Where("ID = ?", movie.ID).Find(&stored)
	assert.Equal(t, movie, stored, "the movie is not stored properly to db")
}

// TestCreateMovieErrorInvalidID creates a movie with invalid id
func TestCreateMovieErrorInvalidID(t *testing.T) {
	// Prepare
	Connect()
	movie := models.Movie{
		ID:   1,
		Name: utils.StrPtr("godfather"),
	}

	// Command
	movie, err := CreateMovie(movie)

	// Assert
	assert.Error(t, err, "invalid id")
}

// TestGetMovieOK gets a movie correctly
func TestGetMovieOK(t *testing.T) {
	// Prepare
	Connect()
	movie := models.Movie{
		Name: utils.StrPtr("godfather"),
	}
	movie, err := CreateMovie(movie)
	assert.NoError(t, err)

	// Command
	response, err := GetMovie(1)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, movie, response, "the movie is not retrieved properly from db")
}

// TestGetMovieErrorNotFound gets a movie that does not exist
func TestGetMovieErrorNotFound(t *testing.T) {
	// Prepare
	Connect()

	// Command
	_, err := GetMovie(1)

	// Assert
	assert.Error(t, err, "record not found")
}

// TestUpdateMovieOK updates a movie correctly
func TestUpdateMovieOK(t *testing.T) {
	// Prepare
	Connect()
	movie := models.Movie{
		Name: utils.StrPtr("godfather"),
	}
	movie, err := CreateMovie(movie)
	assert.NoError(t, err)

	// Command
	movie.Name = utils.StrPtr("godfather 2")
	_, err = UpdateMovie(movie.ID, movie)

	// Assert
	assert.NoError(t, err)
}

// TestUpdateMovieErrorNotFound updates a movie that does not exist
func TestUpdateMovieErrorNotFound(t *testing.T) {
	// Prepare
	Connect()
	movie := models.Movie{
		ID:   1,
		Name: utils.StrPtr("godfather 2"),
	}

	// Command
	_, err := UpdateMovie(movie.ID, movie)

	// Assert
	assert.Error(t, err, "record not found")
}

// TestUpdateMovieErrorInvalidID updates a movie with invalid id
func TestUpdateMovieErrorInvalidID(t *testing.T) {
	// Prepare
	Connect()
	movie := models.Movie{
		Name: utils.StrPtr("godfather"),
	}
	movie, err := CreateMovie(movie)
	assert.NoError(t, err)

	// Command
	movie.ID = 2
	_, err = UpdateMovie(1, movie)

	// Assert
	assert.Error(t, err, "invalid id")
}

// TestListMoviesOK lists movies
func TestListMoviesOK(t *testing.T) {
	// Prepare
	Connect()
	movie := models.Movie{
		Name: utils.StrPtr("godfather"),
	}

	// movies list is empty
	movies := ListMovies()
	assert.Len(t, movies, 0)

	// create a movie
	movie, err := CreateMovie(movie)
	assert.NoError(t, err)

	// movies list has a movie
	movies = ListMovies()
	assert.Len(t, movies, 1)
	assert.Equal(t, movie, *movies[0])
}

// TestDeleteMovieOK deletes a movie
func TestDeleteMovieOK(t *testing.T) {
	// setup
	Connect()
	movie := models.Movie{
		Name: utils.StrPtr("godfather"),
	}
	movie, err := CreateMovie(movie)
	assert.NoError(t, err)

	// execute
	DeleteMovie(movie.ID)
	assert.NoError(t, err)

	// assert
	movie, err = GetMovie(1)
	assert.Error(t, err, "record not found")
}
