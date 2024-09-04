package customerrors

import "errors"

var (
	ErrPasswordIncorrect = errors.New("password is incorrect")
	ErrUserNotFound      = errors.New("user not found")
)
