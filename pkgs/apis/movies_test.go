package apis

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/harbur/golang-gin-starter/pkgs/models"
	"github.com/harbur/golang-gin-starter/pkgs/store"
	"github.com/stretchr/testify/assert"
)

func TestListMoviesWithFuncOK(t *testing.T) {
	// prepare
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// test
	ListMovies(c)

	// assert
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, `[]`, w.Body.String())
}

func TestPostMovieWithRouterOK(t *testing.T) {
	// prepare
	router := gin.New()
	router.POST("/api/movies", PostMovie)

	body := &models.Movie{
		Name: "godfather",
	}
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(body)

	// test
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/movies", buf)
	router.ServeHTTP(w, req)

	// assert
	assert.Equal(t, 201, w.Code)
	assert.Regexp(t, `{"id":.*,"name":"godfather"}`, w.Body.String())
}

func TestPostMovieWithRouterErrorInvalidID(t *testing.T) {
	// prepare
	router := gin.New()
	router.POST("/api/movies", PostMovie)

	body := &models.Movie{
		ID:   1,
		Name: "godfather",
	}
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(body)

	// test
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/movies", buf)
	router.ServeHTTP(w, req)

	// assert
	assert.Equal(t, 409, w.Code)
	assert.Equal(t, `{"error":"invalid id"}`, w.Body.String())
}

func TestPostMovieWithRouterErrorNameIsRequired(t *testing.T) {
	// prepare
	router := gin.New()
	router.POST("/api/movies", PostMovie)

	body := &models.Movie{}
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(body)

	// test
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/movies", buf)
	router.ServeHTTP(w, req)

	// assert
	assert.Equal(t, 400, w.Code)
	assert.Equal(t, `{"error":"Key: 'Movie.Name' Error:Field validation for 'Name' failed on the 'required' tag"}`, w.Body.String())
}

func TestGetMovieWithRouterOK(t *testing.T) {
	// prepare
	router := gin.New()
	router.GET("/api/movies/:id", GetMovie)

	body := models.Movie{
		Name: "godfather",
	}
	_, err := store.Movies.CreateMovie(body)
	assert.NoError(t, err)

	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(body)

	// test
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/movies/1", buf)
	router.ServeHTTP(w, req)

	// assert
	assert.Equal(t, 200, w.Code)
	assert.Regexp(t, `{"id":.*,"name":"godfather"}`, w.Body.String())
}

func TestGetMovieErrorNotFound(t *testing.T) {
	// prepare
	router := gin.New()
	router.GET("/api/movies/:id", GetMovie)

	body := models.Movie{
		Name: "godfather",
	}

	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(body)

	// test
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/movies/999999", buf)
	router.ServeHTTP(w, req)

	// assert
	assert.Equal(t, 404, w.Code)
	assert.Equal(t, `{"error":"record not found"}`, w.Body.String())
}

func TestPutMovieWithRouterOK(t *testing.T) {
	// prepare
	router := gin.New()
	router.PUT("/api/movies/:id", PutMovie)

	body := models.Movie{
		Name: "godfather",
	}
	body, err := store.Movies.CreateMovie(body)
	assert.NoError(t, err)
	body.Name = "godfather 2"

	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(body)

	// test
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/api/movies/"+fmt.Sprint(body.ID), buf)
	router.ServeHTTP(w, req)

	// assert
	assert.Equal(t, 200, w.Code)
	assert.Regexp(t, `{"id":.*,"name":"godfather 2"}`, w.Body.String())
}

func TestPutMovieErrorNotFound(t *testing.T) {
	// prepare
	router := gin.New()
	router.PUT("/api/movies/:id", PutMovie)

	body := models.Movie{
		ID:   99999,
		Name: "godfather 2",
	}

	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(body)

	// test
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/api/movies/99999", buf)
	router.ServeHTTP(w, req)

	// assert
	assert.Equal(t, 404, w.Code)
	assert.Equal(t, `{"error":"record not found"}`, w.Body.String())
}

func TestPutMovieWithRouterErrorNameIsRequired(t *testing.T) {
	// prepare
	router := gin.New()
	router.PUT("/api/movies/:id", PutMovie)

	body := models.Movie{
		ID: 1,
	}

	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(body)

	// test
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/api/movies/1", buf)
	router.ServeHTTP(w, req)

	// assert
	assert.Equal(t, 400, w.Code)
	assert.Equal(t, `{"error":"Key: 'Movie.Name' Error:Field validation for 'Name' failed on the 'required' tag"}`, w.Body.String())
}

func TestDeleteMovieOK(t *testing.T) {
	// prepare
	router := gin.New()
	router.DELETE("/api/movies/:id", DeleteMovie)

	body := models.Movie{
		Name: "godfather",
	}
	body, err := store.Movies.CreateMovie(body)
	assert.NoError(t, err)

	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(body)

	// test
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/api/movies/1", buf)
	router.ServeHTTP(w, req)

	// assert
	assert.Equal(t, 204, w.Code)
	assert.Equal(t, "", w.Body.String())
}
