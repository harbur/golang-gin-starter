package apis

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/harbur/golang-gin-starter/models"
	"github.com/harbur/golang-gin-starter/restapi/operations/healthz"
)

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
func Healthz(params healthz.HealthzParams) middleware.Responder {
	payload := healthcheck(params)
	return healthz.NewHealthzOK().WithPayload(&payload)
}

func healthcheck(params healthz.HealthzParams) models.Healthz {
	res := models.Healthz{
		BuildInfo: &models.BuildInfo{
			BuildDate:  &BuildDate,
			GitBranch:  &GitBranch,
			GitCommit:  &GitCommit,
			GitState:   &GitState,
			GitSummary: &GitSummary,
		},
	}
	return res
}
