// This file is safe to edit. Once it exists it will not be overwritten

package apis

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"bilatung/internal/apis/operations"
	"bilatung/internal/apis/operations/auth"
	"bilatung/internal/apis/operations/health"
	"bilatung/internal/handlers"
	"bilatung/internal/models"
)

//go:generate swagger generate server --target ../../../bilatung --name Bilatung --spec ../../api/swagger.yml --model-package internal/models --server-package internal/apis --principal interface{}

func configureFlags(api *operations.BilatungAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.BilatungAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	// GET HEALTH
	api.HealthGetHealthHandler = health.GetHealthHandlerFunc(func(params health.GetHealthParams) middleware.Responder {
		result, err := handlers.NewHandler().GetHealtcheck()
		if err != nil {
			var errorMessage = new(string)
			*errorMessage = err.Error()
			return health.NewGetHealthDefault(400).WithPayload(&models.Error{Code: "400", Message: *errorMessage})
		}

		return health.NewGetHealthOK().WithPayload(&health.GetHealthOKBody{
			Data: &models.Health{
				Status: result.Status,
			},
			Message: "Success",
		})
	})

	if api.AuthPostLoginHandler == nil {
		api.AuthPostLoginHandler = auth.PostLoginHandlerFunc(func(params auth.PostLoginParams) middleware.Responder {
			return middleware.NotImplemented("operation auth.PostLogin has not yet been implemented")
		})
	}
	if api.AuthPostLogoutHandler == nil {
		api.AuthPostLogoutHandler = auth.PostLogoutHandlerFunc(func(params auth.PostLogoutParams) middleware.Responder {
			return middleware.NotImplemented("operation auth.PostLogout has not yet been implemented")
		})
	}
	if api.AuthPostRegisterHandler == nil {
		api.AuthPostRegisterHandler = auth.PostRegisterHandlerFunc(func(params auth.PostRegisterParams) middleware.Responder {
			return middleware.NotImplemented("operation auth.PostRegister has not yet been implemented")
		})
	}

	api.PreServerShutdown = func() {}

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
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
