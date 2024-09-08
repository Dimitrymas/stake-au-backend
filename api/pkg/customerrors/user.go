package customerrors

import "errors"

var (
	ErrSubNotActive      = errors.New("subscription is not active")
	ErrPasswordIncorrect = errors.New("password is incorrect")
	ErrUserNotFound      = errors.New("user not found")
	ErrUserAlreadyExists = errors.New("user already exists")
	ErrInvalidMnemonic   = errors.New("invalid mnemonic")
)
