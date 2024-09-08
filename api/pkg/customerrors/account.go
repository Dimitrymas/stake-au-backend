package customerrors

import (
	"errors"
	"fmt"
)

var (
	ErrAccountsLimit         = errors.New("maximum number of accounts reached")
	ErrCreatePartialAccounts = errors.New("partial accounts creation")
	ErrAccountNotFound       = errors.New("account not found")
)

type partialAccountsError struct {
	BaseErr    error
	Created    int
	NotCreated int
}

func NewPartialAccountsError(created, notCreated int) error {
	return &partialAccountsError{
		BaseErr:    ErrCreatePartialAccounts,
		Created:    created,
		NotCreated: notCreated,
	}
}

func (e *partialAccountsError) Error() string {
	return fmt.Sprintf("Accounts limit reached. Created: %d, Not created: %d", e.Created, e.NotCreated)
}

func (e *partialAccountsError) Unwrap() error {
	return e.BaseErr
}
