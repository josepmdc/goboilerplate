package json

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/josepmdc/goboilerplate/http/response"
	"github.com/josepmdc/goboilerplate/logger"
)

const contentTypeHeader = `application/json`

// OK creates a new JSON response with a 200 status code.
func OK(w http.ResponseWriter, r *http.Request, body interface{}) {
	builder := response.New(w, r)
	builder.WithHeader("Content-Type", contentTypeHeader)
	builder.WithBody(toJSON(body))
	builder.Write()
}

// Created sends a created response to the client.
func Created(w http.ResponseWriter, r *http.Request, body interface{}) {
	builder := response.New(w, r)
	builder.WithStatus(http.StatusCreated)
	builder.WithHeader("Content-Type", contentTypeHeader)
	builder.WithBody(toJSON(body))
	builder.Write()
}

// NoContent sends a no content response to the client.
func NoContent(w http.ResponseWriter, r *http.Request) {
	builder := response.New(w, r)
	builder.WithStatus(http.StatusNoContent)
	builder.WithHeader("Content-Type", contentTypeHeader)
	builder.Write()
}

// ServerError sends an internal error to the client.
func ServerError(w http.ResponseWriter, r *http.Request, err error) {
	logger.Logger.Warnf("[HTTP:Internal Server Error] %s => %v", r.URL, err)

	builder := response.New(w, r)
	builder.WithStatus(http.StatusInternalServerError)
	builder.WithHeader("Content-Type", contentTypeHeader)
	builder.WithBody(toJSONError(err))
	builder.Write()
}

// BadRequest sends a bad request error to the client.
func BadRequest(w http.ResponseWriter, r *http.Request, err error) {
	logger.Logger.Warnf("[HTTP:Bad Request] %s => %v", r.URL, err)

	builder := response.New(w, r)
	builder.WithStatus(http.StatusBadRequest)
	builder.WithHeader("Content-Type", contentTypeHeader)
	builder.WithBody(toJSONError(err))
	builder.Write()
}

// Unauthorized sends a not authorized error to the client.
func Unauthorized(w http.ResponseWriter, r *http.Request) {
	logger.Logger.Warnf("[HTTP:Unauthorized] %s", r.URL)

	builder := response.New(w, r)
	builder.WithStatus(http.StatusUnauthorized)
	builder.WithHeader("Content-Type", contentTypeHeader)
	builder.WithBody(toJSONError(errors.New("Access Unauthorized")))
	builder.Write()
}

// Forbidden sends a forbidden error to the client.
func Forbidden(w http.ResponseWriter, r *http.Request) {
	logger.Logger.Warnf("[HTTP:Forbidden] %s", r.URL)

	builder := response.New(w, r)
	builder.WithStatus(http.StatusForbidden)
	builder.WithHeader("Content-Type", contentTypeHeader)
	builder.WithBody(toJSONError(errors.New("Access Forbidden")))
	builder.Write()
}

// NotFound sends a page not found error to the client.
func NotFound(w http.ResponseWriter, r *http.Request, entity string) {
	logger.Logger.Warnf("[HTTP:Not Found] %s", r.URL)

	builder := response.New(w, r)
	builder.WithStatus(http.StatusNotFound)
	builder.WithHeader("Content-Type", contentTypeHeader)
	builder.WithBody(toJSONError(fmt.Errorf("%s Not Found", entity)))
	builder.Write()
}

func toJSONError(err error) []byte {
	type errorMsg struct {
		ErrorMessage string `json:"error_message"`
	}

	return toJSON(errorMsg{ErrorMessage: err.Error()})
}

func toJSON(v interface{}) []byte {
	b, err := json.Marshal(v)
	if err != nil {
		logger.Logger.Warnf("[HTTP:JSON] %v", err)
		return []byte("")
	}

	return b
}
