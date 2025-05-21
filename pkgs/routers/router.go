package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/harbur/golang-gin-starter/pkgs/apis"
	"github.com/harbur/golang-gin-starter/pkgs/clients/prometheus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	ginprometheus "github.com/zsais/go-gin-prometheus"
)

// SetupRouter handles the setup of Gin router
func SetupRouter() *gin.Engine {
	// Creates a router without any middleware by default
	r := gin.Default()
	r.SetTrustedProxies([]string{"localhost"})

	// Add Prometheus metrics
	p := ginprometheus.NewPrometheus("gin")
	p.Use(r)
	prometheus.Register()

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())
	r = SetupEndpoints(r)

	return r
}

// SetupEndpoints setups the public endpoints
func SetupEndpoints(r *gin.Engine) *gin.Engine {
	// swagger UI
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/", apis.RedirectTo(307, "/swagger/index.html"))

	// healthz
	r.GET("/healthz", apis.Healthz)
	r.GET("/api/healthz", apis.Healthz)

	// movies
	group := r.Group("/api/movies")
	group.GET("", apis.ListMovies)
	group.POST("", apis.PostMovie)
	group.GET(":id", apis.GetMovie)
	group.PUT(":id", apis.PutMovie)
	group.DELETE(":id", apis.DeleteMovie)

	return r
}
