package json

import (
	"encoding/json"
	"net/http"

	"github.com/josepmdc/goboilerplate/log"
)

const contentTypeHeader = `application/json`

type errorMsg struct {
	ErrorMessage string `json:"error_message"`
}

func writeResponse(w http.ResponseWriter, statusCode int, body interface{}) {
	w.Header().Set("Content-Type", contentTypeHeader)

	response, err := json.Marshal(body)
	if err != nil {
		log.Logger.Warnf("Could not encode JSON response -> %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(statusCode)
	_, err = w.Write(response)
	if err != nil {
		log.Logger.Warnf("Could not write response -> %s", err.Error())
	}
}

// OK creates a new JSON response with a 200 status code.
func OK(w http.ResponseWriter, body interface{}) {
	writeResponse(w, http.StatusOK, body)
}

// Created sends a created response to the client.
func Created(w http.ResponseWriter, body interface{}) {
	writeResponse(w, http.StatusCreated, body)
}

// NoContent sends a no content response to the client.
func NoContent(w http.ResponseWriter) {
	w.Header().Set("Content-Type", contentTypeHeader)
	w.WriteHeader(http.StatusNoContent)
}

// ServerError sends an internal error to the client.
func ServerError(w http.ResponseWriter, err error) {
	writeResponse(w, http.StatusInternalServerError, errorMsg{ErrorMessage: err.Error()})
}

// BadRequest sends a bad request error to the client.
func BadRequest(w http.ResponseWriter, err error) {
	writeResponse(w, http.StatusBadRequest, errorMsg{ErrorMessage: err.Error()})
}

// Unauthorized sends a not authorized error to the client.
func Unauthorized(w http.ResponseWriter) {
	writeResponse(w, http.StatusUnauthorized, errorMsg{ErrorMessage: "Access Unauthorized"})
}

// Forbidden sends a forbidden error to the client.
func Forbidden(w http.ResponseWriter) {
	writeResponse(w, http.StatusForbidden, errorMsg{ErrorMessage: "Access Forbidden"})
}

// NotFound sends a page not found error to the client.
func NotFound(w http.ResponseWriter, entity string) {
	writeResponse(w, http.StatusNotFound, errorMsg{ErrorMessage: entity + " Not Found"})
}
