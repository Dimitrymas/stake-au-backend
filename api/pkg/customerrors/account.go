package customerrors

import "errors"

var (
	ErrSubNotActive  = errors.New("subscription is not active")
	ErrAccountsLimit = errors.New("maximum number of accounts reached")
)
