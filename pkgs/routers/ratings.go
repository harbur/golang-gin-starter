package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/harbur/golang-gin-starter/pkgs/apis"
)

// SetupRouterRatings setups the router for ratings API endpoints
func SetupRouterRatings(r *gin.RouterGroup) *gin.RouterGroup {
	group := r.Group("ratings")
	group.GET("", apis.GetRatings)
	group.POST("", apis.PostRating)
	group.GET(":ratingID", apis.GetRating)
	group.PUT(":ratingID", apis.PutRating)
	group.DELETE(":ratingID", apis.DeleteRating)
	return r
}
