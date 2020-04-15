package apis

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/harbur/golang-gin-starter/pkgs/models"
)

// SetupRouterHealthz setups the router for healthz API endpoints
func SetupRouterHealthz(r *gin.Engine) *gin.Engine {
	group := r.Group("/healthz")
	group.GET("", Healthz)
	return r
}

// GitBranch The branch used to build the binary
var GitBranch string

// GitCommit The commit ID used to build the binary
var GitCommit string

// GitSummary The Git Summary
var GitSummary string

// GitState The Git State (dirty or clean)
var GitState string

// BuildDate The date the binary was built
var BuildDate string

// Healthz handles healthz
func Healthz(c *gin.Context) {
	c.JSON(http.StatusOK, healthcheck())
}

func healthcheck() models.Healthz {
	return models.Healthz{
		BuildInfo: &models.BuildInfo{
			BuildDate:  &BuildDate,
			GitBranch:  &GitBranch,
			GitCommit:  &GitCommit,
			GitState:   &GitState,
			GitSummary: &GitSummary,
		},
	}
}
