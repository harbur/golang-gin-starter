package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/harbur/golang-gin-starter/pkgs/apis"
)

// SetupRouterHealthz setups the router for healthz API endpoints
func SetupRouterHealthz(r *gin.Engine) *gin.Engine {
	r.GET("/healthz", apis.Healthz)
	r.GET("/api/healthz", apis.Healthz)
	return r
}
