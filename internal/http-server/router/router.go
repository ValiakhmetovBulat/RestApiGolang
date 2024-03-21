package router

import (
	mwLogger "RestApiGolang/internal/http-server/middleware/logger"
	log "RestApiGolang/internal/logger"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"net/http"
)

type Method string

const (
	POST   Method = "POST"
	GET    Method = "GET"
	PUT    Method = "PUT"
	DELETE Method = "DELETE"
)

type Route struct {
	Method      Method
	Path        string
	HandlerFunc http.HandlerFunc
}

func NewRouter(routes []Route) *chi.Mux {
	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
	}))

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(mwLogger.New(log.Logger))

	for _, route := range routes {
		r.Method(string(route.Method), route.Path, route.HandlerFunc)
		log.Infof("Route added: %#v", route)
	}

	return r
}
