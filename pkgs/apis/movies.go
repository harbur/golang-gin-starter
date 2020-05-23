package apis

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/harbur/golang-gin-starter/pkgs/clients/prometheus"
	"github.com/harbur/golang-gin-starter/pkgs/models"
	"github.com/harbur/golang-gin-starter/pkgs/store"
	"github.com/harbur/golang-gin-starter/pkgs/utils"
	log "github.com/sirupsen/logrus"
)

// ListMovies lists movies
// @Summary lists movies
// @Description lists movies
// @Tags movies
// @Accept json
// @Produce json
// @Success 200 {string} string	"ok"
// @Router /movies [get]
func ListMovies(c *gin.Context) {
	log.Info("list movies")
	response := store.ListMovies()

	prometheus.IncrementAction("list movies")
	c.JSON(http.StatusOK, response)
}

// PostMovie posts a movie
// @Summary posts a movie
// @Description posts a movie
// @Tags movies
// @Accept json
// @Produce json
// @Param movie body models.Movie true "Movie"
// @Success 200 {string} string	"ok"
// @Router /movies [post]
func PostMovie(c *gin.Context) {
	log.Info("post movie")
	var movie models.Movie
	if err := c.BindJSON(&movie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := store.CreateMovie(movie)
	if err != nil {
		utils.ErrorHandler(c, err)
		return
	}

	c.JSON(http.StatusCreated, response)
}

// GetMovie gets a movie
// @Summary gets a movie
// @Description gets a movie
// @Tags movies
// @Accept json
// @Produce json
// @Param   id     path    int     true        "ID"
// @Success 200 {string} string	"ok"
// @Router /movies/{id} [get]
func GetMovie(c *gin.Context) {
	log.Info("get movie")
	id, _ := strconv.ParseInt(c.Params.ByName("id"), 0, 64)
	response, err := store.GetMovie(uint(id))
	if err != nil {
		utils.ErrorHandler(c, err)
		return
	}

	c.JSON(http.StatusOK, response)
}

// PutMovie puts a movie
// @Summary puts a movie
// @Description puts a movie
// @Tags movies
// @Accept json
// @Produce json
// @Param   id     path    int     true        "ID"
// @Param movie body models.Movie true "Movie"
// @Success 200 {string} string	"ok"
// @Router /movies/{id} [put]
func PutMovie(c *gin.Context) {
	log.Info("put movie")
	id, _ := strconv.ParseInt(c.Params.ByName("id"), 0, 64)
	var movie models.Movie
	if err := c.BindJSON(&movie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := store.UpdateMovie(uint(id), movie)
	if err != nil {
		log.Info(err.Error())
		utils.ErrorHandler(c, err)
		return
	}

	c.JSON(http.StatusOK, response)
}

// DeleteMovie deletes a movie
// @Summary deletes a movie
// @Description deletes a movie
// @Tags movies
// @Accept json
// @Produce json
// @Param   id     path    int     true        "ID"
// @Success 204 {string} string	"ok"
// @Router /movies/{id} [delete]
func DeleteMovie(c *gin.Context) {
	log.Info("delete movie")
	id, _ := strconv.ParseInt(c.Params.ByName("id"), 0, 64)
	store.DeleteMovie(uint(id))

	c.Status(http.StatusNoContent)
}
