package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/harbur/golang-gin-starter/pkgs/apis"
)

// SetupRouterMovies setups the router for movies API endpoints
func SetupRouterMovies(r *gin.Engine) *gin.Engine {
	group := r.Group("/api/movies")
	group.GET("", apis.GetMovies)
	group.POST("", apis.PostMovie)
	group.GET(":id", apis.GetMovie)
	group.PUT(":id", apis.PutMovie)
	group.DELETE(":id", apis.DeleteMovie)
	return r
}
