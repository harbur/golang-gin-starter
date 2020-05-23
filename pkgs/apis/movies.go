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

// ListMovies list movies
// @Summary list movies
// @Description list movies
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

// PostMovie post movie
// @Summary post movie
// @Description post movie
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

	prometheus.IncrementAction("post movie")
	c.JSON(http.StatusCreated, response)
}

// GetMovie get movie
// @Summary get movie
// @Description get movie
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
	prometheus.IncrementAction("get movie")
	c.JSON(http.StatusOK, response)
}

// PutMovie put movie
// @Summary put movie
// @Description put movie
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

	prometheus.IncrementAction("put movie")
	c.JSON(http.StatusOK, response)
}

// DeleteMovie delete movie
// @Summary delete movie
// @Description delete movie
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

	prometheus.IncrementAction("delete movie")
	c.Status(http.StatusNoContent)
}
