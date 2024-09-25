package customerrors

import "errors"

var (
	ErrActivationNotFound = errors.New("activation not found")
)
