package api

import (
	"fmt"
	"net/http"
	"time"

	v1 "github.com/josepmdc/goboilerplate/api/v1"
	"github.com/josepmdc/goboilerplate/app"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// NewRouter creates an http router fot the server
func NewRouter(services *app.Services) http.Handler {
	router := chi.NewRouter()
	// Set up our middleware with sane defaults
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(60 * time.Second))

	// Set up root handlers
	router.Get("/", ping)

	// Set up API V1
	router.Mount("/api/v1/", v1.NewRouter(services))

	return router
}

func ping(w http.ResponseWriter, _ *http.Request) {
	_, err := fmt.Fprintf(w, "Hello Friend")

	if err != nil {
		panic(err)
	}
}
