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

func TestGetMoviesOK(t *testing.T) {
	store.Connect()
	router := SetupRouterMovies(gin.New())

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/movies", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "[]", w.Body.String())
}

func TestPostMovieOK(t *testing.T) {
	store.Connect()
	router := SetupRouterMovies(gin.New())

	body := &models.Movie{
		Name: utils.StrPtr("godfather"),
	}
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(body)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/movies", buf)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, `{"id":1,"name":"godfather"}`, w.Body.String())
}
