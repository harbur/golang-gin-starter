package apis

import (
	"testing"

	"github.com/harbur/golang-starter/models"
	"github.com/harbur/golang-starter/pkgs/store"
	"github.com/harbur/golang-starter/pkgs/utils"
	"github.com/harbur/golang-starter/restapi/operations/movies"
)

func TestGetMoviesOK(t *testing.T) {
	store.Connect()
	GetMovies(movies.GetMoviesParams{})
}

func TestPostMovieOK(t *testing.T) {
	store.Connect()
	PostMovie(movies.PostMovieParams{
		Movie: &models.Movie{
			Name: utils.StrPtr("godfather"),
		},
	})
}

func TestPostMovieErrorInvalidID(t *testing.T) {
	store.Connect()
	PostMovie(movies.PostMovieParams{
		Movie: &models.Movie{
			ID:   1,
			Name: utils.StrPtr("godfather"),
		},
	})
}

func TestGetMovieOK(t *testing.T) {
	store.Connect()
	PostMovie(movies.PostMovieParams{
		Movie: &models.Movie{
			Name: utils.StrPtr("godfather"),
		},
	})

	GetMovie(movies.GetMovieParams{ID: 1})
}

func TestGetMovieErrorNotFound(t *testing.T) {
	store.Connect()
	GetMovie(movies.GetMovieParams{ID: 1})
}

func TestPutMovieOK(t *testing.T) {
	store.Connect()
	PostMovie(movies.PostMovieParams{
		Movie: &models.Movie{
			Name: utils.StrPtr("godfather"),
		},
	})

	PutMovie(movies.PutMovieParams{
		ID: 1,
		Movie: &models.Movie{
			ID:   1,
			Name: utils.StrPtr("godfather 2"),
		},
	})
}

func TestPutMovieErrorNotFound(t *testing.T) {
	store.Connect()
	PutMovie(movies.PutMovieParams{
		ID: 1,
		Movie: &models.Movie{
			ID:   1,
			Name: utils.StrPtr("godfather 2"),
		},
	})
}

func TestDeleteMovieOK(t *testing.T) {
	store.Connect()
	PostMovie(movies.PostMovieParams{
		Movie: &models.Movie{
			Name: utils.StrPtr("godfather"),
		},
	})

	DeleteMovie(movies.DeleteMovieParams{
		ID: 1,
	})
}
