package errs

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestNewInternalServerError(t *testing.T) {
	err := NewInternalServerErr("this is the message", errors.New("error"))
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.StatusCode())
	assert.EqualValues(t, "this is the message", err.Message())
	assert.EqualValues(t, "internal_server_error", err.Err())
}

func TestNewAuthenticationError(t *testing.T) {
	err := NewAuthenticationErr("this is the message")
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusUnauthorized, err.StatusCode())
	assert.EqualValues(t, "this is the message", err.Message())
	assert.EqualValues(t, "unauthorized", err.Err())
}

func TestNewNotFoundError(t *testing.T) {
	err := NewNotFoundErr("this is the message")
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode())
	assert.EqualValues(t, "this is the message", err.Message())
	assert.EqualValues(t, "not_found", err.Err())
}

func TestNewBadRequestError(t *testing.T) {
	err := NewBadRequestErr("this is the message")
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusBadRequest, err.StatusCode())
	assert.EqualValues(t, "this is the message", err.Message())
	assert.EqualValues(t, "bad_request", err.Err())
}

func TestNewAuthorizationError(t *testing.T) {
	err := NewAuthorizationErr("this is the message")
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusForbidden, err.StatusCode())
	assert.EqualValues(t, "this is the message", err.Message())
	assert.EqualValues(t, "forbidden", err.Err())
}
