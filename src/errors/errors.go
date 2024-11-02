package errors

import (
	"errors"
)

var (
	ErrBookingAlreadyExists = errors.New("the requested booking cannot be processed as it already exists")
	ErrBookingNotFound      = errors.New("the requested booking was not found")
)

type GenericErrorResponse struct {
	Error string `json:"error"`
}
