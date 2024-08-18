package exception

import (
	"contact-management-restful/models/dto"
	"encoding/json"
	"errors"
	"github.com/go-playground/validator/v10"
	"net/http"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request, err any) {
	if validationError(w, err.(error)) {
		return
	}

	if badRequest(w, err.(error)) {
		return
	}

	if unAuthorized(w, err.(error)) {
		return
	}

	if notFound(w, err.(error)) {
		return
	}

	internalServerError(w, err.(error))
}

func notFound(w http.ResponseWriter, err error) bool {
	var exception NotFound
	ok := errors.As(err, &exception)
	if !ok {
		return false
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)

	webResponse := dto.WebResponse{
		Message: http.StatusText(http.StatusNotFound),
		Error:   err.Error(),
	}

	err = json.NewEncoder(w).Encode(&webResponse)
	if err != nil {
		return false
	}
	return true
}

func unAuthorized(w http.ResponseWriter, err error) bool {
	var exception Unauthorized
	ok := errors.As(err, &exception)
	if !ok {
		return false
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnauthorized)

	webResponse := dto.WebResponse{
		Message: http.StatusText(http.StatusUnauthorized),
		Error:   err.Error(),
	}

	err = json.NewEncoder(w).Encode(&webResponse)
	if err != nil {
		return false
	}
	return true
}

func badRequest(w http.ResponseWriter, err error) bool {
	var exception BadRequest
	ok := errors.As(err, &exception)
	if !ok {
		return false
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)

	webResponse := dto.WebResponse{
		Message: http.StatusText(http.StatusBadRequest),
		Error:   err.Error(),
	}

	err = json.NewEncoder(w).Encode(&webResponse)
	if err != nil {
		return false
	}
	return true
}

func validationError(w http.ResponseWriter, err error) bool {
	var exception validator.ValidationErrors
	ok := errors.As(err, &exception)
	if !ok {
		return false
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)

	webResponse := dto.WebResponse{
		Message: http.StatusText(http.StatusBadRequest),
		Error:   exception.Error(),
	}

	err = json.NewEncoder(w).Encode(&webResponse)
	if err != nil {
		return false
	}
	return true
}

func internalServerError(w http.ResponseWriter, err error) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)

	webResponse := dto.WebResponse{
		Message: http.StatusText(http.StatusInternalServerError),
		Error:   err.Error(),
	}

	_ = json.NewEncoder(w).Encode(&webResponse)
}
