package apis

import (
	"testing"

	"github.com/harbur/golang-gin-starter/restapi/operations/healthz"
)

func TestHealthzOK(t *testing.T) {
	Healthz(healthz.HealthzParams{})
}
