package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/harbur/golang-gin-starter/pkgs/store"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// SetupRouter handles the setup of Gin router
func SetupRouter() *gin.Engine {
	// Creates a router without any middleware by default
	r := gin.Default()

	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	r.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())
	store.Connect()
	r = SetupRouterHealthz(r)
	r = SetupRouterMovies(r)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}
