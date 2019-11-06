package utils

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/harbur/golang-starter/models"
	"github.com/harbur/golang-starter/restapi/operations/movies"
)

// StrPtr string pointer
func StrPtr(s string) *string { return &s }

// BoolPtr boolean pointer
func BoolPtr(s bool) *bool { return &s }

// Int64Ptr int64 pointer
func Int64Ptr(s int64) *int64 { return &s }

// ErrorHandler handles error responses
func ErrorHandler(err error) middleware.Responder {
	if err.Error() == "record not found" {
		return movies.NewGetMovieNotFound().WithPayload(wrapErrorString(err, 404))
	} else if err.Error() == "invalid id" {
		return movies.NewGetMovieConflict().WithPayload(wrapErrorString(err, 409))
	} else {
		return movies.NewGetMovieInternalServerError().WithPayload(wrapErrorString(err, 500))
	}
}

func wrapErrorString(err error, code int64) *models.Error {
	response := &models.Error{
		Message: err.Error(),
		Code:    code,
	}
	return response
}
