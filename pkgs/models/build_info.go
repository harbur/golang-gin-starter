package models

// BuildInfo build info
type BuildInfo struct {

	// build date
	// Required: true
	BuildDate *string `json:"buildDate"`

	// git branch
	// Required: true
	GitBranch *string `json:"gitBranch"`

	// git commit
	// Required: true
	GitCommit *string `json:"gitCommit"`

	// git state
	// Required: true
	GitState *string `json:"gitState"`

	// git summary
	// Required: true
	GitSummary *string `json:"gitSummary"`
}
