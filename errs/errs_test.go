package errs

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestNewInternalServerError(t *testing.T) {
	err := NewInternalServerError("this is the message")
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.StatusCode)
	assert.EqualValues(t, "this is the message", err.Message)
	assert.EqualValues(t, "internal_server_error", err.Error)
}

func TestNewAuthenticationError(t *testing.T) {
	err := NewAuthenticationError("this is the message")
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusUnauthorized, err.StatusCode)
	assert.EqualValues(t, "this is the message", err.Message)
	assert.EqualValues(t, "unauthorized", err.Error)
}

func TestNewNotFoundError(t *testing.T) {
	err := NewNotFoundError("this is the message")
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "this is the message", err.Message)
	assert.EqualValues(t, "not_found", err.Error)
}

func TestNewBadRequestError(t *testing.T) {
	err := NewBadRequestError("this is the message")
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusBadRequest, err.StatusCode)
	assert.EqualValues(t, "this is the message", err.Message)
	assert.EqualValues(t, "bad_request", err.Error)
}

func TestNewAuthorizationError(t *testing.T) {
	err := NewAuthorizationError("this is the message")
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusForbidden, err.StatusCode)
	assert.EqualValues(t, "this is the message", err.Message)
	assert.EqualValues(t, "forbidden", err.Error)
}
