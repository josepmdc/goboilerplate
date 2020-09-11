package v1

import (
	"net/http"

	"github.com/go-chi/chi"
)

type authHandler struct {
	//	service app.AuthService
}

// NewUserHandler creates a handler for the user endpoints
func NewAuthHandler( /*s app.AuthService*/ ) *authHandler {
	return &authHandler{}
}

func (h *authHandler) routes() http.Handler {
	r := chi.NewRouter()

	r.Post("/signin", h.signIn)
	return r
}

func (h *authHandler) signIn(w http.ResponseWriter, r *http.Request) {

}
