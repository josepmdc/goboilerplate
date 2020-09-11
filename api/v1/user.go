package v1

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/josepmdc/goboilerplate/api/v1/models"
	"github.com/josepmdc/goboilerplate/app"
	"github.com/josepmdc/goboilerplate/http/response/json"
	"github.com/josepmdc/goboilerplate/logger"
)

const USER_ID = "userID"

type userHandler struct {
	service app.UserService
}

// NewUserHandler creates a handler for the user endpoints
func NewUserHandler(s app.UserService) *userHandler {
	return &userHandler{
		service: s,
	}
}

func (h *userHandler) routes() http.Handler {
	r := chi.NewRouter()

	r.Get("/{"+USER_ID+"}", h.getUser)
	r.Post("/new", h.createUser)
	r.Post("/check_email", h.checkEmail)
	r.Post("/check_username", h.checkUsername)
	return r
}

func (h *userHandler) getUser(w http.ResponseWriter, r *http.Request) {
	ID, err := uuid.Parse(chi.URLParam(r, USER_ID))
	if err != nil {
		json.BadRequest(w, r, err)
		return
	}

	user, err := h.service.FindUser(ID)
	if err != nil {
		json.NotFound(w, r, "User")
		return
	}
	json.OK(w, r, models.MapUserToAPI(user))
}

func (h *userHandler) createUser(w http.ResponseWriter, r *http.Request) {
	user, err := models.DecodeUser(r.Body)
	if err != nil {
		logger.Logger.Warnf("Could not decode user: %s", err.Error())
		json.BadRequest(w, r, err)
		return
	}
	_, err = h.service.CreateUser(models.MapUserToDomain(user))
	if err != nil {
		logger.Logger.Warnf("Could not create user: %s", err.Error())
		json.BadRequest(w, r, err)
		return
	}

	json.OK(w, r, "User created successfully")
}

func (h *userHandler) checkEmail(w http.ResponseWriter, r *http.Request) {
	credentials, err := models.DecodeCredentials(r.Body)
	if err != nil {
		logger.Logger.Warnf("Could not decode user credentials: %s", err.Error())
		json.BadRequest(w, r, err)
		return
	}
	exists := h.service.CheckEmail(credentials.Email)
	if exists {
		json.BadRequest(w, r, errors.New("Email is already taken"))
		return
	}
	json.OK(w, r, "Email is free")
}

func (h *userHandler) checkUsername(w http.ResponseWriter, r *http.Request) {
	credentials, err := models.DecodeCredentials(r.Body)
	if err != nil {
		logger.Logger.Warnf("Could not decode user credentials: %s", err.Error())
		json.BadRequest(w, r, err)
		return
	}
	exists := h.service.CheckUsername(credentials.UserName)
	if exists {
		json.BadRequest(w, r, errors.New("Username is already taken"))
		return
	}
	json.OK(w, r, "Username is free")
}
