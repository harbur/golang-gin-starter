package apis

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"gotest.tools/assert"
)

func TestRedirectTo(t *testing.T) {
	// prepare
	router := gin.New()
	router.GET("/", RedirectTo(307, "swagger/index.html"))

	// test
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	router.ServeHTTP(w, req)

	// assert
	assert.Equal(t, 307, w.Code)
}
