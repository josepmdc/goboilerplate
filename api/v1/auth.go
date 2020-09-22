package v1

import (
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/josepmdc/goboilerplate/api/v1/models"
	"github.com/josepmdc/goboilerplate/app"
	"github.com/josepmdc/goboilerplate/http/response/json"
	"github.com/josepmdc/goboilerplate/log"
)

type authHandler struct {
	service app.AuthService
}

// NewAuthHandler creates a handler for the authentication endpoints
func NewAuthHandler(s app.AuthService) *authHandler {
	return &authHandler{
		service: s,
	}
}

func (h *authHandler) routes() http.Handler {
	r := chi.NewRouter()

	r.Post("/signin", h.signIn)

	return r
}

func (h *authHandler) signIn(w http.ResponseWriter, r *http.Request) {
	creds, err := models.DecodeCredentials(r.Body)
	if err != nil {
		log.Logger.Warnf("Could not decode credentials: %s", err.Error())
		json.BadRequest(w, r, err)
		return
	}

	token, err := h.service.SignIn(models.MapCredentialsToDomain(creds))
	if err != nil {
		log.Logger.Warnf("Could not sign in: %s", err.Error())
		json.BadRequest(w, r, err)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   token,
		Expires: time.Now().Add(5 * time.Minute),
	})
}
