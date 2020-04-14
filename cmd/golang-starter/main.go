package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/harbur/golang-gin-starter/cmd/golang-starter/docs"
	"github.com/harbur/golang-gin-starter/pkgs/apis"
	"github.com/harbur/golang-gin-starter/pkgs/store"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

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
	r = apis.SetupRouterMovies(r)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}

// @title Golang Starter
// @version 1.0.0
// @description Golang Starter.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.email contact@harbur.io

// @license.name MIT
// @license.url https://github.com/harbur/golang-gin-starter/blob/master/LICENSE

// @BasePath /api/
func main() {
	r := SetupRouter()
	r.Run(fmt.Sprintf(":%v", 8080))

}
