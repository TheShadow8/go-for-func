package util

import (
	"errors"
)

const (
	InvalidEmail       = "invalid email"
	EmailAlreadyExists = "email already exists"
	EmptyPassword      = "password can't be empty"
	InvalidAuthToken   = "invalid auth-token"
	InvalidCredentials = "invalid credentials"
	Unauthorized       = "Unauthorized"
)

var (
	ErrInvalidEmail       = errors.New(InvalidEmail)
	ErrEmailAlreadyExists = errors.New(EmailAlreadyExists)
	ErrEmptyPassword      = errors.New(EmptyPassword)
	ErrInvalidAuthToken   = errors.New(InvalidAuthToken)
	ErrInvalidCredentials = errors.New(InvalidCredentials)
	ErrUnauthorized       = errors.New(Unauthorized)
)

var errorCode = map[string]int{
	InvalidEmail:       1,
	EmailAlreadyExists: 2,
	EmptyPassword:      2,
	InvalidAuthToken:   4,
	InvalidCredentials: 5,
	Unauthorized:       6,
}

// StatusText returns a text for the HTTP status code. It returns the empty
// string if the code is unknown.
func GetErrorCode(error error) int {
	code := errorCode[error.Error()]

	if code == 0 {
		code = 500
	}

	return code
}
