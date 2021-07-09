package errs

import (
	"errors"
	"fmt"
	"net/http"
)

type RestErr interface {
	Message() string
	StatusCode() int
	Err() string
	Causes() []interface{}
}

type restErr struct {
	message    string        `json:"message,omitempty"`
	statusCode int           `json:"code,omitempty"`
	err        string        `json:"error,omitempty"`
	causes     []interface{} `json:"causes,omitempty"`
}

func (e restErr) Message() string {
	return e.message
}

func (e restErr) StatusCode() int {
	return e.statusCode
}

func (e restErr) Err() string {
	return e.err
}

func (e restErr) Causes() []interface{} {
	return e.causes
}

func (e restErr) Error() string {
	return fmt.Sprintf("message: %s - status: %d - error: %s - causes: [ %v ]", e.message, e.statusCode, e.err, e.causes)
}

func NewError(msg string) error {
	return errors.New(msg)
}

func NewRestErr(message string, statusCode int, error string, causes []interface{}) RestErr {
	return &restErr{message: message, statusCode: statusCode, err: error, causes: causes}
}

func NewBadRequestErr(message string) RestErr {
	return &restErr{
		message:    message,
		statusCode: http.StatusBadRequest,
		err:        "bad_request",
	}
}

func NewInternalServerErr(message string, err error) RestErr {
	r := &restErr{
		message:    message,
		statusCode: http.StatusInternalServerError,
		err:        "internal_server_error",
		causes:     []interface{}{},
	}
	if err != nil {
		r.causes = append(r.causes, err.Error())
	}
	return r
}

func NewNotFoundErr(message string) RestErr {
	return &restErr{
		message:    message,
		statusCode: http.StatusNotFound,
		err:        "not_found",
	}
}

func NewAuthenticationErr(message string) RestErr {
	return &restErr{
		message:    message,
		statusCode: http.StatusUnauthorized,
		err:        "unauthorized",
	}
}

func NewAuthorizationErr(message string) RestErr {
	return &restErr{
		message:    message,
		statusCode: http.StatusForbidden,
		err:        "forbidden",
	}
}
