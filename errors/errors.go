package errors

import (
	"net/http"
)

type RestErr struct {
	Message    string        `json:"message,omitempty"`
	StatusCode int           `json:"code,omitempty"`
	Error      string        `json:"error,omitempty"`
	Causes     []interface{} `json:"causes"`
}

func NewBadRequestError(message string) *RestErr {
	return &RestErr{
		Message:    message,
		StatusCode: http.StatusBadRequest,
		Error:      "bad_request",
	}
}

func NewInternalServerError(message string, err error) *RestErr {
	restErr := &RestErr{
		Message:    message,
		StatusCode: http.StatusInternalServerError,
		Error:      "internal_server_error",
		Causes:     []interface{}{},
	}
	if err != nil {
		restErr.Causes = append(restErr.Causes, err.Error())
	}
	return restErr
}

func NewNotFoundError(message string) *RestErr {
	return &RestErr{
		Message:    message,
		StatusCode: http.StatusNotFound,
		Error:      "not_found",
	}
}

func NewAuthenticationError(message string) *RestErr {
	return &RestErr{
		Message:    message,
		StatusCode: http.StatusUnauthorized,
		Error:      "unauthorized",
	}
}

func NewAuthorizationError(message string) *RestErr {
	return &RestErr{
		Message:    message,
		StatusCode: http.StatusForbidden,
		Error:      "forbidden",
	}
}
