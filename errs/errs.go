package errs

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type RestErr interface {
	Message() string
	StatusCode() int
	Err() string
	Causes() []string
}

type Err struct {
	Msg         string   `json:"message,omitempty"`
	Code        int      `json:"code,omitempty"`
	ErrMessage  string   `json:"error,omitempty"`
	ErrorCauses []string `json:"causes,omitempty"`
}

func (e Err) Message() string {
	return e.Msg
}

func (e Err) StatusCode() int {
	return e.Code
}

func (e Err) Err() string {
	return e.ErrMessage
}

func (e Err) Causes() []string {
	return e.ErrorCauses
}

func FromBytes(errorBytes []byte) (RestErr, error) {
	var r Err
	err := json.Unmarshal(errorBytes, &r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (e Err) Error() string {
	return fmt.Sprintf("message: %s - status: %d - error: %s - causes: [ %v ]", e.Msg, e.Code, e.ErrMessage, e.ErrorCauses)
}

func NewError(msg string) error {
	return errors.New(msg)
}

func NewRestErr(message string, statusCode int, error string, causes []string) RestErr {
	return &Err{Msg: message, Code: statusCode, ErrMessage: error, ErrorCauses: causes}
}

func NewBadRequestErr(message string) RestErr {
	return &Err{
		Msg:        message,
		Code:       http.StatusBadRequest,
		ErrMessage: "bad_request",
	}
}

func NewInternalServerErr(message string, err error) RestErr {
	r := &Err{
		Msg:         message,
		Code:        http.StatusInternalServerError,
		ErrMessage:  "internal_server_error",
		ErrorCauses: []string{},
	}
	if err != nil {
		r.ErrorCauses = append(r.ErrorCauses, err.Error())
	}
	return r
}

func NewNotFoundErr(message string) RestErr {
	return &Err{
		Msg:        message,
		Code:       http.StatusNotFound,
		ErrMessage: "not_found",
	}
}

func NewAuthenticationErr(message string) RestErr {
	return &Err{
		Msg:        message,
		Code:       http.StatusUnauthorized,
		ErrMessage: "unauthorized",
	}
}

func NewAuthorizationErr(message string) RestErr {
	return &Err{
		Msg:        message,
		Code:       http.StatusForbidden,
		ErrMessage: "forbidden",
	}
}
