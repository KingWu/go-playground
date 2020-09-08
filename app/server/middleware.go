package server

import (
	"github.com/phyber/negroni-gzip/gzip"
	"github.com/rs/cors"
	"github.com/urfave/negroni"
)

func BuildMiddleware(allowedOrigins []string, debug bool) *negroni.Negroni {
	middlewareController := negroni.New()
	middlewareController.Use(negroni.NewRecovery())
	middlewareController.Use(negroni.NewLogger())

	supportGzip(middlewareController)
	supportCors(middlewareController, allowedOrigins, debug)
	
	return middlewareController
}

func supportGzip(middlewareController *negroni.Negroni) {
	middlewareController.Use(gzip.Gzip(gzip.DefaultCompression))
}

func supportCors(middlewareController *negroni.Negroni, allowedOrigins []string, debug bool) {
	cors := cors.New(cors.Options{
		AllowedOrigins: allowedOrigins,
    AllowCredentials: true,
    Debug: debug,
	})
	middlewareController.Use(cors)
}
