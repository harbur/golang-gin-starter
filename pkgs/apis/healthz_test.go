package apis

import (
	"testing"

	"github.com/harbur/golang-starter/restapi/operations/healthz"
)

func TestHealthzOK(t *testing.T) {
	Healthz(healthz.HealthzParams{})
}
