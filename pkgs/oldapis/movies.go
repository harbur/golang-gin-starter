package apis

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/harbur/golang-gin-starter/pkgs/store"
	"github.com/harbur/golang-gin-starter/pkgs/utils"
	"github.com/harbur/golang-gin-starter/restapi/operations/movies"
	log "github.com/sirupsen/logrus"
)

// GetMovies lists movies
func GetMovies(params movies.GetMoviesParams) middleware.Responder {
	log.Info("get movies")
	response := store.ListMovies()
	return movies.NewGetMoviesOK().WithPayload(response)
}

// GetMovie gets a movie
func GetMovie(params movies.GetMovieParams) middleware.Responder {
	log.Info("get movie")
	response, err := store.GetMovie(params.ID)
	if err != nil {
		return utils.ErrorHandler(err)
	}

	return movies.NewGetMovieOK().WithPayload(&response)
}

// PostMovie posts a movie
func PostMovie(params movies.PostMovieParams) middleware.Responder {
	log.Info("post movie")
	response, err := store.CreateMovie(*params.Movie)
	if err != nil {
		return utils.ErrorHandler(err)
	}

	return movies.NewPostMovieCreated().WithPayload(&response)
}

// PutMovie puts a movie
func PutMovie(params movies.PutMovieParams) middleware.Responder {
	log.Info("put movie")
	response, err := store.UpdateMovie(params.ID, *params.Movie)
	if err != nil {
		log.Info(err.Error())
		return utils.ErrorHandler(err)
	}

	return movies.NewPutMovieOK().WithPayload(&response)
}

// DeleteMovie gets a movie
func DeleteMovie(params movies.DeleteMovieParams) middleware.Responder {
	log.Info("delete movie")
	store.DeleteMovie(params.ID)
	return movies.NewDeleteMovieOK()
}
