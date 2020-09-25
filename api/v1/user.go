package v1

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/josepmdc/goboilerplate/api/v1/models"
	"github.com/josepmdc/goboilerplate/app"
	"github.com/josepmdc/goboilerplate/http/response/json"
	"github.com/josepmdc/goboilerplate/log"
)

const userID = "userID"

// UserHandler is a struct that contains all the functions for handling requests
// to the user endpoints
type UserHandler struct {
	service app.UserService
}

// NewUserHandler creates a handler for the user endpoints and sets the necessary
// values for it to operate
func NewUserHandler(s app.UserService) *UserHandler {
	return &UserHandler{
		service: s,
	}
}

func (h *UserHandler) routes() http.Handler {
	r := chi.NewRouter()

	r.Get("/{"+userID+"}", h.getUser)
	r.Post("/new", h.createUser)
	r.Post("/check_email", h.checkEmail)
	r.Post("/check_username", h.checkUsername)

	return r
}

func (h *UserHandler) getUser(w http.ResponseWriter, r *http.Request) {
	ID, err := uuid.Parse(chi.URLParam(r, userID))
	if err != nil {
		log.Logger.Warnf("[HTTP:Bad Request] %s => %v", r.URL, err)
		json.BadRequest(w, err)
		return
	}

	user, err := h.service.FindByID(ID)
	if err != nil {
		json.NotFound(w, "User")
		return
	}

	json.OK(w, models.MapUserToAPI(user))
}

func (h *UserHandler) createUser(w http.ResponseWriter, r *http.Request) {
	credentials, err := models.DecodeCredentials(r.Body)
	if err != nil {
		log.Logger.Warnf("Could not decode user: %s", err.Error())
		json.BadRequest(w, err)
		return
	}

	_, err = h.service.CreateUser(models.MapCredentialsToDomain(credentials))
	if err != nil {
		log.Logger.Warnf("Could not create user: %s", err.Error())
		json.BadRequest(w, err)
		return
	}

	json.OK(w, "User created successfully")
}

func (h *UserHandler) checkEmail(w http.ResponseWriter, r *http.Request) {
	credentials, err := models.DecodeCredentials(r.Body)
	if err != nil {
		log.Logger.Warnf("Could not decode user credentials: %s", err.Error())
		json.BadRequest(w, err)
		return
	}

	exists := h.service.CheckEmail(credentials.Email)
	if exists {
		json.BadRequest(w, errors.New("Email is already taken"))
		return
	}

	json.OK(w, "Email is free")
}

func (h *UserHandler) checkUsername(w http.ResponseWriter, r *http.Request) {
	credentials, err := models.DecodeCredentials(r.Body)
	if err != nil {
		log.Logger.Warnf("Could not decode user credentials: %s", err.Error())
		json.BadRequest(w, err)
		return
	}

	exists := h.service.CheckUsername(credentials.Username)
	if exists {
		json.BadRequest(w, errors.New("Username is already taken"))
		return
	}

	json.OK(w, "Username is free")
}
