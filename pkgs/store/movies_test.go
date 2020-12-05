package store

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/harbur/golang-gin-starter/pkgs/models"
)

// TestCreateMovieOK creates a movie correctly
func TestCreateMovieOK(t *testing.T) {
	// Prepare
	movie := models.Movie{
		Name: "godfather",
	}

	// Command
	movie, err := Movies.CreateMovie(movie)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, uint(1), movie.ID, "the movie ID should now be 1")
	stored, err := Movies.GetMovie(movie.ID)
	assert.NotNil(t, stored)
	assert.NoError(t, err)
	assert.Equal(t, movie.Name, stored.Movie.Name, "the movie is not stored properly to db")
}

// TestCreateMovieErrorInvalidID creates a movie with invalid id
func TestCreateMovieErrorInvalidID(t *testing.T) {
	// Prepare
	movie := models.Movie{
		ID:   1,
		Name: "godfather",
	}

	// Command
	movie, err := Movies.CreateMovie(movie)

	// Assert
	assert.Error(t, err, "invalid id")
}

// TestGetMovieOK gets a movie correctly
func TestGetMovieOK(t *testing.T) {
	// Prepare
	movie := models.Movie{
		Name: "godfather",
	}
	movie, err := Movies.CreateMovie(movie)
	assert.NoError(t, err)

	// Command
	response, err := Movies.GetMovie(1)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, movie.Name, response.Movie.Name, "the movie is not retrieved properly from db")
}

// TestGetMovieErrorNotFound gets a movie that does not exist
func TestGetMovieErrorNotFound(t *testing.T) {
	// Command
	_, err := Movies.GetMovie(99999)

	// Assert
	assert.Error(t, err, "record not found")
}

// TestUpdateMovieOK updates a movie correctly
func TestUpdateMovieOK(t *testing.T) {
	// Prepare
	movie := models.Movie{
		Name: "godfather",
	}
	movie, err := Movies.CreateMovie(movie)
	assert.NoError(t, err)

	// Command
	movie.Name = "godfather 2"
	_, err = Movies.UpdateMovie(movie.ID, movie)

	// Assert
	assert.NoError(t, err)
}

// TestUpdateMovieErrorNotFound updates a movie that does not exist
func TestUpdateMovieErrorNotFound(t *testing.T) {
	// Prepare
	movie := models.Movie{
		ID:   99999,
		Name: "godfather 2",
	}

	// Command
	_, err := Movies.UpdateMovie(movie.ID, movie)

	// Assert
	assert.Error(t, err, "record not found")
}

// TestUpdateMovieErrorInvalidID updates a movie with invalid id
func TestUpdateMovieErrorInvalidID(t *testing.T) {
	// Prepare
	movie := models.Movie{
		Name: "godfather",
	}
	movie, err := Movies.CreateMovie(movie)
	assert.NoError(t, err)

	// Command
	movie.ID = 2
	_, err = Movies.UpdateMovie(1, movie)

	// Assert
	assert.Error(t, err, "invalid id")
}

// TestListMoviesOK lists movies
func TestListMoviesOK(t *testing.T) {
	// Prepare
	movie := models.Movie{
		Name: "godfather",
	}

	// movies list is empty
	movies := Movies.ListMovies()
	length := len(movies.Movies)

	// create a movie
	movie, err := Movies.CreateMovie(movie)
	assert.NoError(t, err)

	// movies list has a movie
	movies = Movies.ListMovies()
	assert.Len(t, movies.Movies, length+1)
	assert.Equal(t, movie.Name, movies.Movies[0].Name)
}

// TestDeleteMovieOK deletes a movie
func TestDeleteMovieOK(t *testing.T) {
	// setup
	movie := models.Movie{
		Name: "godfather",
	}
	movie, err := Movies.CreateMovie(movie)
	assert.NoError(t, err)

	// execute
	Movies.DeleteMovie(movie.ID)
	assert.NoError(t, err)

	// assert
	_, err = Movies.GetMovie(movie.ID)
	assert.Error(t, err, "record not found")
}
