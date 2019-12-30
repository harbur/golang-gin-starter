// Code generated by go-swagger; DO NOT EDIT.

package movies

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetMovieHandlerFunc turns a function with the right signature into a get movie handler
type GetMovieHandlerFunc func(GetMovieParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetMovieHandlerFunc) Handle(params GetMovieParams) middleware.Responder {
	return fn(params)
}

// GetMovieHandler interface for that can handle valid get movie params
type GetMovieHandler interface {
	Handle(GetMovieParams) middleware.Responder
}

// NewGetMovie creates a new http.Handler for the get movie operation
func NewGetMovie(ctx *middleware.Context, handler GetMovieHandler) *GetMovie {
	return &GetMovie{Context: ctx, Handler: handler}
}

/*GetMovie swagger:route GET /api/movies/{id} movies getMovie

Get movie

Gets a movie.

*/
type GetMovie struct {
	Context *middleware.Context
	Handler GetMovieHandler
}

func (o *GetMovie) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetMovieParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
