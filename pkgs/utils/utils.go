package utils

import (
	"github.com/gin-gonic/gin"
)

// StrPtr string pointer
func StrPtr(s string) *string { return &s }

// BoolPtr boolean pointer
func BoolPtr(s bool) *bool { return &s }

// Int64Ptr int64 pointer
func Int64Ptr(s int64) *int64 { return &s }

// ErrorHandler handles error responses
func ErrorHandler(c *gin.Context, err error) {
	if err.Error() == "record not found" {
		wrapErrorString(c, err, 404)
	} else if err.Error() == "invalid id" {
		wrapErrorString(c, err, 409)
	} else {
		wrapErrorString(c, err, 500)
	}
}

func wrapErrorString(c *gin.Context, err error, code int) {
	content := gin.H{"error": err.Error()}
	c.JSON(code, content)
}
