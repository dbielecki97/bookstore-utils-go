package errors

import (
	"net/http"
)

type RestErr struct {
	Message    string        `json:"message,omitempty"`
	StatusCode int           `json:"code,omitempty"`
	Error      string        `json:"error,omitempty"`
	Causes     []interface{} `json:"causes,omitempty"`
}

func NewBadRequestError(message string) *RestErr {
	return &RestErr{
		Message:    message,
		StatusCode: http.StatusBadRequest,
		Error:      "bad_request",
	}
}

func NewInternalServerError(message string, err error) *RestErr {
	r := &RestErr{
		Message:    message,
		StatusCode: http.StatusInternalServerError,
		Error:      "internal_server_error",
		Causes:     []interface{}{},
	}
	if err != nil {
		r.Causes = append(r.Causes, err.Error())
	}
	return r
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
