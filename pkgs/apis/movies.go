package apis

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/harbur/golang-gin-starter/models"
	"github.com/harbur/golang-gin-starter/pkgs/store"
	"github.com/harbur/golang-gin-starter/pkgs/utils"
	"github.com/harbur/golang-gin-starter/restapi/operations/movies"
	log "github.com/sirupsen/logrus"
)

// GetMovies lists movies
// @Summary lists movies
// @Description lists movies
// @Accept json
// @Produce json
// @Success 200 {string} string	"ok"
// @Router /movies [get]
func GetMovies(c *gin.Context) {
	log.Info("get movies")
	response := store.ListMovies()
	c.JSON(http.StatusOK, response)
}

// GetMovie gets a movie
// @Summary gets a movie
// @Description gets a movie
// @Accept json
// @Produce json
// @Param   id     path    int     true        "ID"
// @Success 200 {string} string	"ok"
// @Router /movies/{id} [get]
func GetMovie(c *gin.Context) {
	log.Info("get movie")
	id, _ := strconv.ParseInt(c.Params.ByName("id"), 0, 64)
	response, err := store.GetMovie(id)
	if err != nil {
		utils.ErrorHandler(c, err)
		return
	}

	c.JSON(http.StatusOK, response)
}

// PostMovie posts a movie
// @Summary posts a movie
// @Description posts a movie
// @Accept json
// @Produce json
// @Success 200 {string} string	"ok"
// @Router /movies/{id} [post]
func PostMovie(c *gin.Context) {
	log.Info("post movie")
	var movie models.Movie
	if err := c.ShouldBindJSON(&movie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := store.CreateMovie(movie)
	if err != nil {
		utils.ErrorHandler(c, err)
		return
	}

	c.JSON(http.StatusOK, response)
}

// PutMovie puts a movie
func PutMovie(c *gin.Context, params movies.PutMovieParams) {
	log.Info("put movie")
	response, err := store.UpdateMovie(params.ID, *params.Movie)
	if err != nil {
		log.Info(err.Error())
		utils.ErrorHandler(c, err)
		return
	}

	c.JSON(http.StatusOK, response)
}

// DeleteMovie gets a movie
func DeleteMovie(c *gin.Context, params movies.DeleteMovieParams) {
	log.Info("delete movie")
	store.DeleteMovie(params.ID)

	c.JSON(http.StatusOK, nil)
}
