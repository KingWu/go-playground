package server

import (
	"github.com/phyber/negroni-gzip/gzip"
	"github.com/rs/cors"
	"github.com/urfave/negroni"
)

func BuildMiddleware() *negroni.Negroni {
	middlewareController := negroni.New()
	setUpMiddleware(middlewareController)
	return middlewareController
}

func setUpMiddleware(middlewareController *negroni.Negroni) {
	middlewareController.Use(negroni.NewRecovery())
	middlewareController.Use(negroni.NewLogger())

	supportGzip(middlewareController)
	supportCors(middlewareController)
}

func supportGzip(middlewareController *negroni.Negroni) {
	middlewareController.Use(gzip.Gzip(gzip.DefaultCompression))
}

func supportCors(middlewareController *negroni.Negroni) {
	cors := cors.New(cors.Options{
		AllowedOrigins: []string{"http://foo.com", "http://foo.com:8080"},
    AllowCredentials: true,
    // Enable Debugging for testing, consider disabling in production
    Debug: true,
	})
	middlewareController.Use(cors)
}
