// Code generated by go-swagger; DO NOT EDIT.

package movies

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/harbur/golang-gin-starter/models"
)

// GetMoviesOKCode is the HTTP code returned for type GetMoviesOK
const GetMoviesOKCode int = 200

/*GetMoviesOK OK

swagger:response getMoviesOK
*/
type GetMoviesOK struct {

	/*
	  In: Body
	*/
	Payload models.Movies `json:"body,omitempty"`
}

// NewGetMoviesOK creates GetMoviesOK with default headers values
func NewGetMoviesOK() *GetMoviesOK {

	return &GetMoviesOK{}
}

// WithPayload adds the payload to the get movies o k response
func (o *GetMoviesOK) WithPayload(payload models.Movies) *GetMoviesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get movies o k response
func (o *GetMoviesOK) SetPayload(payload models.Movies) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetMoviesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = models.Movies{}
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetMoviesConflictCode is the HTTP code returned for type GetMoviesConflict
const GetMoviesConflictCode int = 409

/*GetMoviesConflict Conflict

swagger:response getMoviesConflict
*/
type GetMoviesConflict struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetMoviesConflict creates GetMoviesConflict with default headers values
func NewGetMoviesConflict() *GetMoviesConflict {

	return &GetMoviesConflict{}
}

// WithPayload adds the payload to the get movies conflict response
func (o *GetMoviesConflict) WithPayload(payload *models.Error) *GetMoviesConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get movies conflict response
func (o *GetMoviesConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetMoviesConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetMoviesInternalServerErrorCode is the HTTP code returned for type GetMoviesInternalServerError
const GetMoviesInternalServerErrorCode int = 500

/*GetMoviesInternalServerError Internal Server Error

swagger:response getMoviesInternalServerError
*/
type GetMoviesInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetMoviesInternalServerError creates GetMoviesInternalServerError with default headers values
func NewGetMoviesInternalServerError() *GetMoviesInternalServerError {

	return &GetMoviesInternalServerError{}
}

// WithPayload adds the payload to the get movies internal server error response
func (o *GetMoviesInternalServerError) WithPayload(payload *models.Error) *GetMoviesInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get movies internal server error response
func (o *GetMoviesInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetMoviesInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
