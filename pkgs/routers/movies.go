package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/harbur/golang-gin-starter/pkgs/apis"
)

// SetupRouterMovies setups the router for movies API endpoints
func SetupRouterMovies(r *gin.RouterGroup) *gin.RouterGroup {
	movies := r.Group("movies")
	movies.GET("", apis.GetMovies)
	movies.POST("", apis.PostMovie)
	movie := movies.Group(":movieID")
	movie.GET("", apis.GetMovie)
	movie.PUT("", apis.PutMovie)
	movie.DELETE("", apis.DeleteMovie)
	movie = SetupRouterRatings(movie)
	return r
}
