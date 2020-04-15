package store

import "github.com/harbur/golang-gin-starter/pkgs/models"

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

func Healthz() models.Healthz {
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
