package errors

import (
	"errors"
)

var (
	ErrBookingAlreadyExists = errors.New("the requested booking cannot be processed as it already exists")
)
