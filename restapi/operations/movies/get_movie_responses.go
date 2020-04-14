// Code generated by go-swagger; DO NOT EDIT.

package movies

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/harbur/golang-gin-starter/models"
)

// GetMovieOKCode is the HTTP code returned for type GetMovieOK
const GetMovieOKCode int = 200

/*GetMovieOK OK

swagger:response getMovieOK
*/
type GetMovieOK struct {

	/*
	  In: Body
	*/
	Payload *models.Movie `json:"body,omitempty"`
}

// NewGetMovieOK creates GetMovieOK with default headers values
func NewGetMovieOK() *GetMovieOK {

	return &GetMovieOK{}
}

// WithPayload adds the payload to the get movie o k response
func (o *GetMovieOK) WithPayload(payload *models.Movie) *GetMovieOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get movie o k response
func (o *GetMovieOK) SetPayload(payload *models.Movie) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetMovieOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetMovieNotFoundCode is the HTTP code returned for type GetMovieNotFound
const GetMovieNotFoundCode int = 404

/*GetMovieNotFound Not Found

swagger:response getMovieNotFound
*/
type GetMovieNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetMovieNotFound creates GetMovieNotFound with default headers values
func NewGetMovieNotFound() *GetMovieNotFound {

	return &GetMovieNotFound{}
}

// WithPayload adds the payload to the get movie not found response
func (o *GetMovieNotFound) WithPayload(payload *models.Error) *GetMovieNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get movie not found response
func (o *GetMovieNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetMovieNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetMovieConflictCode is the HTTP code returned for type GetMovieConflict
const GetMovieConflictCode int = 409

/*GetMovieConflict Conflict

swagger:response getMovieConflict
*/
type GetMovieConflict struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetMovieConflict creates GetMovieConflict with default headers values
func NewGetMovieConflict() *GetMovieConflict {

	return &GetMovieConflict{}
}

// WithPayload adds the payload to the get movie conflict response
func (o *GetMovieConflict) WithPayload(payload *models.Error) *GetMovieConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get movie conflict response
func (o *GetMovieConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetMovieConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetMovieInternalServerErrorCode is the HTTP code returned for type GetMovieInternalServerError
const GetMovieInternalServerErrorCode int = 500

/*GetMovieInternalServerError Internal Server Error

swagger:response getMovieInternalServerError
*/
type GetMovieInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetMovieInternalServerError creates GetMovieInternalServerError with default headers values
func NewGetMovieInternalServerError() *GetMovieInternalServerError {

	return &GetMovieInternalServerError{}
}

// WithPayload adds the payload to the get movie internal server error response
func (o *GetMovieInternalServerError) WithPayload(payload *models.Error) *GetMovieInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get movie internal server error response
func (o *GetMovieInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetMovieInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
