package domain

import (
	"fmt"
	"net/http"
)

type AppError struct {
	HTTPStatus int    `json:"-"`
	Code       string `json:"Code,omitempty"`
	Message    string `json:"message,omitempty"`
}

func (e AppError) Error() string {
	return fmt.Sprintf("%s - %s", e.Code, e.Message)
}

// NewAppError New functions create a new AppError instance
func NewAppError(httpStatus int, label, message string) *AppError {
	return &AppError{
		HTTPStatus: httpStatus,
		Code:       label,
		Message:    message,
	}
}

var (
	ErrInternal      = &AppError{HTTPStatus: http.StatusInternalServerError, Code: "INTERNAL", Message: "Internal error"}
	ErrNotFound      = &AppError{HTTPStatus: http.StatusNotFound, Code: "NOT_FOUND", Message: "Resource not found"}
	ErrStale         = &AppError{HTTPStatus: http.StatusBadRequest, Code: "STALE", Message: "Resource is stale.  please retry"}
	ErrAlreadyExists = &AppError{HTTPStatus: http.StatusBadRequest, Code: "ALREADY_EXISTS", Message: "Resource already exists"}
)
