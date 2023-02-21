package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

// Decode the JSON body from a POST request into a struct
func decodeJSONBody(w http.ResponseWriter, r *http.Request, m interface{}) *Error {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	errResponse := Error{
		Status:    http.StatusBadRequest,
		Error:     http.StatusText(http.StatusBadRequest),
		Timestamp: time.Now(),
	}
	// Struct error handling
	if err := decoder.Decode(&m); err != nil {
		var syntaxErr *json.SyntaxError
		var typeErr *json.UnmarshalTypeError
		switch {
		case errors.As(err, &syntaxErr):
			errResponse.Message = fmt.Sprintf("Request body contains badly-formed JSON (at position %d)", syntaxErr.Offset)
		case errors.Is(err, io.ErrUnexpectedEOF):
			errResponse.Message = "Request body contains badly-formed JSON"
		case errors.As(err, &typeErr):
			errResponse.Message = fmt.Sprintf(
				"Request body contains an invalid value for the '%q' field (at position %d)", typeErr.Field, typeErr.Offset,
			)
		case errors.Is(err, io.EOF):
			errResponse.Message = "Request body must not be empty"
		case strings.HasPrefix(err.Error(), "json: unknown field "):
			field := strings.TrimPrefix(err.Error(), "json: unknown field ")
			errResponse.Message = fmt.Sprintf("Request body contains unknown field '%s'", field)
		default:
			errResponse.Message = "Unable to parse request body"
		}
		return &errResponse
	}
	return nil
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	errResponse := Error{
		Status:    http.StatusNotFound,
		Error:     http.StatusText(http.StatusNotFound),
		Message:   "Endpoint requested was not found",
		Timestamp: time.Now(),
	}
	res, _ := json.Marshal(errResponse)
	w.Write(res)
}

func methodNotAllowed(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusMethodNotAllowed)
	errResponse := Error{
		Status:    http.StatusMethodNotAllowed,
		Error:     http.StatusText(http.StatusMethodNotAllowed),
		Message:   fmt.Sprintf("Endpoint does not allow %s method", r.Method),
		Timestamp: time.Now(),
	}
	res, _ := json.Marshal(errResponse)
	w.Write(res)
}

func emptySlug(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadRequest)
	errResponse := Error{
		Status:    http.StatusBadRequest,
		Error:     http.StatusText(http.StatusBadRequest),
		Message:   "Song slug path parameter cannot be empty",
		Timestamp: time.Now(),
	}
	res, _ := json.Marshal(errResponse)
	w.Write(res)
}
