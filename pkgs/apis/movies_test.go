package apis

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/harbur/golang-gin-starter/pkgs/models"
	"github.com/harbur/golang-gin-starter/pkgs/store"
	"github.com/harbur/golang-gin-starter/pkgs/utils"
	"gotest.tools/assert"
)

func TestGetMoviesWithFuncOK(t *testing.T) {
	// prepare
	store.Connect()
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// test
	GetMovies(c)

	// assert
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "[]", w.Body.String())
}

func TestPostMovieWithRouterOK(t *testing.T) {
	// prepare
	store.Connect()
	router := gin.New()
	router.POST("/api/movies", PostMovie)

	body := &models.Movie{
		Name: utils.StrPtr("godfather"),
	}
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(body)

	// test
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/movies", buf)
	router.ServeHTTP(w, req)

	// assert
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, `{"id":1,"name":"godfather"}`, w.Body.String())
}

func TestPostMovieWithRouterErrorInvalidID(t *testing.T) {
	// prepare
	store.Connect()
	router := gin.New()
	router.POST("/api/movies", PostMovie)

	body := &models.Movie{
		ID:   1,
		Name: utils.StrPtr("godfather"),
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

func TestGetMovieWithRouterOK(t *testing.T) {
	// prepare
	store.Connect()
	router := gin.New()
	router.GET("/api/movies/:id", GetMovie)

	body := models.Movie{
		Name: utils.StrPtr("godfather"),
	}
	_, err := store.CreateMovie(body)
	assert.NilError(t, err)

	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(body)

	// test
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/movies/1", buf)
	router.ServeHTTP(w, req)

	// assert
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, `{"id":1,"name":"godfather"}`, w.Body.String())
}

func TestGetMovieErrorNotFound(t *testing.T) {
	// prepare
	store.Connect()
	router := gin.New()
	router.GET("/api/movies/:id", GetMovie)

	body := models.Movie{
		Name: utils.StrPtr("godfather"),
	}

	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(body)

	// test
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/movies/1", buf)
	router.ServeHTTP(w, req)

	// assert
	assert.Equal(t, 404, w.Code)
	assert.Equal(t, `{"error":"record not found"}`, w.Body.String())
}

func TestPutMovieWithRouterOK(t *testing.T) {
	// prepare
	store.Connect()
	router := gin.New()
	router.PUT("/api/movies/:id", PutMovie)

	body := models.Movie{
		Name: utils.StrPtr("godfather"),
	}
	body, err := store.CreateMovie(body)
	assert.NilError(t, err)
	body.Name = utils.StrPtr("godfather 2")

	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(body)

	// test
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/api/movies/1", buf)
	router.ServeHTTP(w, req)

	// assert
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, `{"id":1,"name":"godfather 2"}`, w.Body.String())
}

func TestPutMovieErrorNotFound(t *testing.T) {
	// prepare
	store.Connect()
	router := gin.New()
	router.PUT("/api/movies/:id", PutMovie)

	body := models.Movie{
		ID:   1,
		Name: utils.StrPtr("godfather 2"),
	}

	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(body)

	// test
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/api/movies/1", buf)
	router.ServeHTTP(w, req)

	// assert
	assert.Equal(t, 404, w.Code)
	assert.Equal(t, `{"error":"record not found"}`, w.Body.String())
}
func TestDeleteMovieOK(t *testing.T) {
	// prepare
	store.Connect()
	router := gin.New()
	router.DELETE("/api/movies/:id", DeleteMovie)

	body := models.Movie{
		Name: utils.StrPtr("godfather"),
	}
	body, err := store.CreateMovie(body)
	assert.NilError(t, err)

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
