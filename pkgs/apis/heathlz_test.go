package apis

import (
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"gotest.tools/assert"
)

func TestHealthzOK(t *testing.T) {
	// prepare
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// test
	Healthz(c)

	// assert
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, `{"buildInfo":{"buildDate":"","gitBranch":"","gitCommit":"","gitState":"","gitSummary":""}}`, w.Body.String())
}
