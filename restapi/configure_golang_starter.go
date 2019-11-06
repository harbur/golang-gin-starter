// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"

	"github.com/harbur/golang-starter/pkgs/apis"
	"github.com/harbur/golang-starter/pkgs/middlewares"
	"github.com/harbur/golang-starter/pkgs/store"
	"github.com/harbur/golang-starter/restapi/operations"
	"github.com/harbur/golang-starter/restapi/operations/healthz"
	"github.com/harbur/golang-starter/restapi/operations/movies"
)

//go:generate swagger generate server --target ../../golang-starter --name GolangStarter --spec ../swagger.yaml

func configureFlags(api *operations.GolangStarterAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.GolangStarterAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()
	store.Connect()

	// movies
	api.MoviesGetMoviesHandler = movies.GetMoviesHandlerFunc(apis.GetMovies)
	api.MoviesPostMovieHandler = movies.PostMovieHandlerFunc(apis.PostMovie)
	api.MoviesGetMovieHandler = movies.GetMovieHandlerFunc(apis.GetMovie)
	api.MoviesPutMovieHandler = movies.PutMovieHandlerFunc(apis.PutMovie)
	api.MoviesDeleteMovieHandler = movies.DeleteMovieHandlerFunc(apis.DeleteMovie)

	// healthz
	api.HealthzHealthzHandler = healthz.HealthzHandlerFunc(apis.Healthz)

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return middlewares.SwaggerUIMiddleware(handler)
}
