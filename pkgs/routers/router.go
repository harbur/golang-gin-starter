package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/harbur/golang-gin-starter/pkgs/apis"
	"github.com/harbur/golang-gin-starter/pkgs/clients/prometheus"
	"github.com/harbur/golang-gin-starter/pkgs/store"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	ginprometheus "github.com/zsais/go-gin-prometheus"
)

// SetupRouter handles the setup of Gin router
func SetupRouter() *gin.Engine {
	// Creates a router without any middleware by default
	r := gin.Default()

	// Add Prometheus metrics
	p := ginprometheus.NewPrometheus("gin")
	p.Use(r)
	prometheus.Register()

	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	r.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())
	store.Connect()
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
