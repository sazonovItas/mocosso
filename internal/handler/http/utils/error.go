package httputils

import (
	"encoding/json"
	"net/http"
)

const InternalErrorMessage = "Internal server error"

// Error defines model for Error.
type Error struct {
	// Code Error code
	Code int `json:"code"`

	// Message Error message
	Message string `json:"message"`

	// Description Error description
	Description string `json:"description,omitempty"`
}

func NewError(code int, message string, description string) Error {
	return Error{
		Code:        code,
		Message:     message,
		Description: description,
	}
}

func NewInternalError(description string) Error {
	return Error{
		Code:        http.StatusInternalServerError,
		Message:     InternalErrorMessage,
		Description: description,
	}
}

func (err Error) WriteResponse(w http.ResponseWriter) {
	w.WriteHeader(err.Code)
	w.Header().Add("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(err)
}

func DefaultErrorHandlerFunc(w http.ResponseWriter, r *http.Request, err error) {
	NewError(http.StatusBadRequest, err.Error(), err.Error()).WriteResponse(w)
}
