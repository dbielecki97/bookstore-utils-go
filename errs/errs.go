package errs

import (
	"net/http"
)

type RestErr struct {
	Message    string        `json:"message,omitempty"`
	StatusCode int           `json:"code,omitempty"`
	Error      string        `json:"error,omitempty"`
	Causes     []interface{} `json:"causes,omitempty"`
}

func NewBadRequestErr(message string) *RestErr {
	return &RestErr{
		Message:    message,
		StatusCode: http.StatusBadRequest,
		Error:      "bad_request",
	}
}

func NewInternalServerErr(message string, err error) *RestErr {
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

func NewNotFoundErr(message string) *RestErr {
	return &RestErr{
		Message:    message,
		StatusCode: http.StatusNotFound,
		Error:      "not_found",
	}
}

func NewAuthenticationErr(message string) *RestErr {
	return &RestErr{
		Message:    message,
		StatusCode: http.StatusUnauthorized,
		Error:      "unauthorized",
	}
}

func NewAuthorizationErr(message string) *RestErr {
	return &RestErr{
		Message:    message,
		StatusCode: http.StatusForbidden,
		Error:      "forbidden",
	}
}
