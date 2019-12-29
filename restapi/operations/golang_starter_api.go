// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	errors "github.com/go-openapi/errors"
	loads "github.com/go-openapi/loads"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	security "github.com/go-openapi/runtime/security"
	spec "github.com/go-openapi/spec"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/harbur/golang-starter/restapi/operations/healthz"
	"github.com/harbur/golang-starter/restapi/operations/movies"
)

// NewGolangStarterAPI creates a new GolangStarter instance
func NewGolangStarterAPI(spec *loads.Document) *GolangStarterAPI {
	return &GolangStarterAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		customConsumers:     make(map[string]runtime.Consumer),
		customProducers:     make(map[string]runtime.Producer),
		ServerShutdown:      func() {},
		spec:                spec,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,
		JSONConsumer:        runtime.JSONConsumer(),
		JSONProducer:        runtime.JSONProducer(),
		MoviesDeleteMovieHandler: movies.DeleteMovieHandlerFunc(func(params movies.DeleteMovieParams) middleware.Responder {
			return middleware.NotImplemented("operation MoviesDeleteMovie has not yet been implemented")
		}),
		MoviesGetMovieHandler: movies.GetMovieHandlerFunc(func(params movies.GetMovieParams) middleware.Responder {
			return middleware.NotImplemented("operation MoviesGetMovie has not yet been implemented")
		}),
		MoviesGetMoviesHandler: movies.GetMoviesHandlerFunc(func(params movies.GetMoviesParams) middleware.Responder {
			return middleware.NotImplemented("operation MoviesGetMovies has not yet been implemented")
		}),
		HealthzHealthzHandler: healthz.HealthzHandlerFunc(func(params healthz.HealthzParams) middleware.Responder {
			return middleware.NotImplemented("operation HealthzHealthz has not yet been implemented")
		}),
		MoviesPostMovieHandler: movies.PostMovieHandlerFunc(func(params movies.PostMovieParams) middleware.Responder {
			return middleware.NotImplemented("operation MoviesPostMovie has not yet been implemented")
		}),
		MoviesPutMovieHandler: movies.PutMovieHandlerFunc(func(params movies.PutMovieParams) middleware.Responder {
			return middleware.NotImplemented("operation MoviesPutMovie has not yet been implemented")
		}),
	}
}

/*GolangStarterAPI Golang Starter */
type GolangStarterAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	customConsumers map[string]runtime.Consumer
	customProducers map[string]runtime.Producer
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator
	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator
	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for a "application/json" mime type
	JSONConsumer runtime.Consumer

	// JSONProducer registers a producer for a "application/json" mime type
	JSONProducer runtime.Producer

	// MoviesDeleteMovieHandler sets the operation handler for the delete movie operation
	MoviesDeleteMovieHandler movies.DeleteMovieHandler
	// MoviesGetMovieHandler sets the operation handler for the get movie operation
	MoviesGetMovieHandler movies.GetMovieHandler
	// MoviesGetMoviesHandler sets the operation handler for the get movies operation
	MoviesGetMoviesHandler movies.GetMoviesHandler
	// HealthzHealthzHandler sets the operation handler for the healthz operation
	HealthzHealthzHandler healthz.HealthzHandler
	// MoviesPostMovieHandler sets the operation handler for the post movie operation
	MoviesPostMovieHandler movies.PostMovieHandler
	// MoviesPutMovieHandler sets the operation handler for the put movie operation
	MoviesPutMovieHandler movies.PutMovieHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// SetDefaultProduces sets the default produces media type
func (o *GolangStarterAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *GolangStarterAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *GolangStarterAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *GolangStarterAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *GolangStarterAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *GolangStarterAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *GolangStarterAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the GolangStarterAPI
func (o *GolangStarterAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.MoviesDeleteMovieHandler == nil {
		unregistered = append(unregistered, "movies.DeleteMovieHandler")
	}

	if o.MoviesGetMovieHandler == nil {
		unregistered = append(unregistered, "movies.GetMovieHandler")
	}

	if o.MoviesGetMoviesHandler == nil {
		unregistered = append(unregistered, "movies.GetMoviesHandler")
	}

	if o.HealthzHealthzHandler == nil {
		unregistered = append(unregistered, "healthz.HealthzHandler")
	}

	if o.MoviesPostMovieHandler == nil {
		unregistered = append(unregistered, "movies.PostMovieHandler")
	}

	if o.MoviesPutMovieHandler == nil {
		unregistered = append(unregistered, "movies.PutMovieHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *GolangStarterAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *GolangStarterAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {

	return nil

}

// Authorizer returns the registered authorizer
func (o *GolangStarterAPI) Authorizer() runtime.Authorizer {

	return nil

}

// ConsumersFor gets the consumers for the specified media types
func (o *GolangStarterAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {

	result := make(map[string]runtime.Consumer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONConsumer

		}

		if c, ok := o.customConsumers[mt]; ok {
			result[mt] = c
		}
	}
	return result

}

// ProducersFor gets the producers for the specified media types
func (o *GolangStarterAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {

	result := make(map[string]runtime.Producer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONProducer

		}

		if p, ok := o.customProducers[mt]; ok {
			result[mt] = p
		}
	}
	return result

}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *GolangStarterAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the golang starter API
func (o *GolangStarterAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *GolangStarterAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened

	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/api/movies/{id}"] = movies.NewDeleteMovie(o.context, o.MoviesDeleteMovieHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/api/movies/{id}"] = movies.NewGetMovie(o.context, o.MoviesGetMovieHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/api/movies"] = movies.NewGetMovies(o.context, o.MoviesGetMoviesHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/api/healthz"] = healthz.NewHealthz(o.context, o.HealthzHealthzHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/api/movies"] = movies.NewPostMovie(o.context, o.MoviesPostMovieHandler)

	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/api/movies/{id}"] = movies.NewPutMovie(o.context, o.MoviesPutMovieHandler)

}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *GolangStarterAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *GolangStarterAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *GolangStarterAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *GolangStarterAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}