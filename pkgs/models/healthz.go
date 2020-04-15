package models

// Healthz healthz
type Healthz struct {

	// build info
	// Required: true
	BuildInfo *BuildInfo `json:"buildInfo"`
}
