package utils

import (
	"errors"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestStrPtr(t *testing.T) {
	StrPtr("test")
}

func TestBoolPtr(t *testing.T) {
	BoolPtr(true)
}

func TestInt64Ptr(t *testing.T) {
	Int64Ptr(1)
}

func TestErrorHandler(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	err := errors.New("hello")
	ErrorHandler(c, err)
}
