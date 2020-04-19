package apis

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/harbur/golang-gin-starter/pkgs/models"
	"github.com/harbur/golang-gin-starter/pkgs/store"
	"github.com/harbur/golang-gin-starter/pkgs/utils"
	log "github.com/sirupsen/logrus"
)

// GetRatings lists ratings
// @Summary lists ratings
// @Description lists ratings
// @Tags ratings
// @Accept json
// @Produce json
// @Param   movieID  path    int     true        "Movie ID"
// @Success 200 {string} string	"ok"
// @Router /movies/{movieID}/ratings [get]
func GetRatings(c *gin.Context) {
	log.Info("get ratings")
	movieID, _ := strconv.ParseInt(c.Params.ByName("movieID"), 0, 64)
	response := store.ListRatings(uint(movieID))
	c.JSON(http.StatusOK, response)
}

// PostRating posts a rating
// @Summary posts a rating
// @Description posts a rating
// @Tags ratings
// @Accept json
// @Produce json
// @Param   movieID  path    int     true        "Movie ID"
// @Param   rating   body models.Rating true "Rating"
// @Success 200 {string} string	"ok"
// @Router /movies/{movieID}/ratings [post]
func PostRating(c *gin.Context) {
	log.Info("post rating")
	var rating models.Rating
	if err := c.BindJSON(&rating); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := store.CreateRating(rating)
	if err != nil {
		utils.ErrorHandler(c, err)
		return
	}

	c.JSON(http.StatusOK, response)
}

// GetRating gets a rating
// @Summary gets a rating
// @Description gets a rating
// @Tags ratings
// @Accept json
// @Produce json
// @Param   movieID  path    int     true        "Movie ID"
// @Param   id       path    int     true        "ID"
// @Success 200 {string} string	"ok"
// @Router /movies/{movieID}/ratings/{id} [get]
func GetRating(c *gin.Context) {
	log.Info("get rating")
	id, _ := strconv.ParseInt(c.Params.ByName("id"), 0, 64)
	response, err := store.GetRating(uint(id))
	if err != nil {
		utils.ErrorHandler(c, err)
		return
	}

	c.JSON(http.StatusOK, response)
}

// PutRating puts a rating
// @Summary puts a rating
// @Description puts a rating
// @Tags ratings
// @Accept json
// @Produce json
// @Param   movieID  path    int     true        "Movie ID"
// @Param   id       path    int     true        "ID"
// @Param rating body models.Rating true "Rating"
// @Success 200 {string} string	"ok"
// @Router /movies/{movieID}/ratings/{id} [put]
func PutRating(c *gin.Context) {
	log.Info("put rating")
	id, _ := strconv.ParseInt(c.Params.ByName("id"), 0, 64)
	var rating models.Rating
	if err := c.BindJSON(&rating); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := store.UpdateRating(uint(id), rating)
	if err != nil {
		log.Info(err.Error())
		utils.ErrorHandler(c, err)
		return
	}

	c.JSON(http.StatusOK, response)
}

// DeleteRating deletes a rating
// @Summary deletes a rating
// @Description deletes a rating
// @Tags ratings
// @Accept json
// @Produce json
// @Param   movieID  path    int     true        "Movie ID"
// @Param   id       path    int     true        "ID"
// @Success 204 {string} string	"ok"
// @Router /movies/{movieID}/ratings/{id} [delete]
func DeleteRating(c *gin.Context) {
	log.Info("delete rating")
	id, _ := strconv.ParseInt(c.Params.ByName("id"), 0, 64)
	store.DeleteRating(uint(id))

	c.Status(http.StatusNoContent)
}
