package v1

import (
	"net/http"

	"github.com/josepmdc/goboilerplate/app"

	"github.com/go-chi/chi"
)

// NewRouter creates an http router for the API
func NewRouter(s *app.Services) http.Handler {
	r := chi.NewRouter()
	r.Mount("/user", NewUserHandler(s.User).routes())
	r.Mount("/auth", NewAuthHandler(s.Auth).routes())
	return r
}
