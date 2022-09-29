package lib

import "errors"

var (
	ErrSecretIsEmpty    = errors.New("secret is empty")
	ErrNotValidPassword = errors.New("not valid password")
	ErrNotValidEmail    = errors.New("not valid email string")
	ErrDontShowCount    = errors.New("you can't watch the secret anymore")
	ErrEmptyParameter   = errors.New("empty parameter")
)
