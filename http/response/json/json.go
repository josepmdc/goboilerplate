package json

import (
	"encoding/json"
	"net/http"
)

const contentTypeHeader = `application/json`

type errorMsg struct {
	ErrorMessage string `json:"error_message"`
}

// OK creates a new JSON response with a 200 status code.
func OK(w http.ResponseWriter, body interface{}) {
	w.Header().Set("Content-Type", contentTypeHeader)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(body)
}

// Created sends a created response to the client.
func Created(w http.ResponseWriter, body interface{}) {
	w.Header().Set("Content-Type", contentTypeHeader)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(body)
}

// NoContent sends a no content response to the client.
func NoContent(w http.ResponseWriter) {
	w.Header().Set("Content-Type", contentTypeHeader)
	w.WriteHeader(http.StatusNoContent)
}

// ServerError sends an internal error to the client.
func ServerError(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", contentTypeHeader)
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(errorMsg{ErrorMessage: err.Error()})
}

// BadRequest sends a bad request error to the client.
func BadRequest(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", contentTypeHeader)
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(errorMsg{ErrorMessage: err.Error()})
}

// Unauthorized sends a not authorized error to the client.
func Unauthorized(w http.ResponseWriter) {
	w.Header().Set("Content-Type", contentTypeHeader)
	w.WriteHeader(http.StatusUnauthorized)
	json.NewEncoder(w).Encode(errorMsg{ErrorMessage: "Access Unauthorized"})
}

// Forbidden sends a forbidden error to the client.
func Forbidden(w http.ResponseWriter) {
	w.Header().Set("Content-Type", contentTypeHeader)
	w.WriteHeader(http.StatusForbidden)
	json.NewEncoder(w).Encode(errorMsg{ErrorMessage: "Access Forbidden"})
}

// NotFound sends a page not found error to the client.
func NotFound(w http.ResponseWriter, entity string) {
	w.Header().Set("Content-Type", contentTypeHeader)
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(errorMsg{ErrorMessage: entity + " Not Found"})
}
