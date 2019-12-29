// Code generated by go-swagger; DO NOT EDIT.

package movies

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/harbur/golang-starter/models"
)

// PostMovieOKCode is the HTTP code returned for type PostMovieOK
const PostMovieOKCode int = 200

/*PostMovieOK OK

swagger:response postMovieOK
*/
type PostMovieOK struct {

	/*
	  In: Body
	*/
	Payload *models.Movie `json:"body,omitempty"`
}

// NewPostMovieOK creates PostMovieOK with default headers values
func NewPostMovieOK() *PostMovieOK {

	return &PostMovieOK{}
}

// WithPayload adds the payload to the post movie o k response
func (o *PostMovieOK) WithPayload(payload *models.Movie) *PostMovieOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post movie o k response
func (o *PostMovieOK) SetPayload(payload *models.Movie) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostMovieOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostMovieConflictCode is the HTTP code returned for type PostMovieConflict
const PostMovieConflictCode int = 409

/*PostMovieConflict Conflict

swagger:response postMovieConflict
*/
type PostMovieConflict struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostMovieConflict creates PostMovieConflict with default headers values
func NewPostMovieConflict() *PostMovieConflict {

	return &PostMovieConflict{}
}

// WithPayload adds the payload to the post movie conflict response
func (o *PostMovieConflict) WithPayload(payload *models.Error) *PostMovieConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post movie conflict response
func (o *PostMovieConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostMovieConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostMovieInternalServerErrorCode is the HTTP code returned for type PostMovieInternalServerError
const PostMovieInternalServerErrorCode int = 500

/*PostMovieInternalServerError Internal Server Error

swagger:response postMovieInternalServerError
*/
type PostMovieInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostMovieInternalServerError creates PostMovieInternalServerError with default headers values
func NewPostMovieInternalServerError() *PostMovieInternalServerError {

	return &PostMovieInternalServerError{}
}

// WithPayload adds the payload to the post movie internal server error response
func (o *PostMovieInternalServerError) WithPayload(payload *models.Error) *PostMovieInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post movie internal server error response
func (o *PostMovieInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostMovieInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}