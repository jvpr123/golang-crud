package rest_error

import "net/http"

type ErrorCause struct {
	Field   string
	Message string
}

type RestError struct {
	Message   string       `json:"message"`
	ErrorName string       `json:"error"`
	ErrorCode int          `json:"code"`
	Causes    []ErrorCause `json:"causes,omitempty"`
}

func (r *RestError) Error() string {
	return r.Message
}

func NewRestError(message string, errorName string, errorCode int, causes []ErrorCause) *RestError {
	return &RestError{
		Message:   message,
		ErrorName: errorName,
		ErrorCode: errorCode,
		Causes:    causes,
	}
}

func NewBadRequestError(message string) *RestError {
	return &RestError{
		Message:   message,
		ErrorName: "Bad Request",
		ErrorCode: http.StatusBadRequest,
	}
}

func NewBadRequestValidationError(message string, causes []ErrorCause) *RestError {
	return &RestError{
		Message:   message,
		ErrorName: "Bad Request",
		ErrorCode: http.StatusBadRequest,
		Causes:    causes,
	}
}

func NewForbiddenError(message string) *RestError {
	return &RestError{
		Message:   message,
		ErrorName: "Forbidden",
		ErrorCode: http.StatusForbidden,
	}
}

func NewNotFoundError(message string) *RestError {
	return &RestError{
		Message:   message,
		ErrorName: "Not Found",
		ErrorCode: http.StatusNotFound,
	}
}

func NewInternalServerError(message string) *RestError {
	return &RestError{
		Message:   message,
		ErrorName: "Internal Server Error",
		ErrorCode: http.StatusInternalServerError,
	}
}
