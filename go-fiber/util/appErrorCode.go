package util

import (
	"errors"
)

const (
	InvalidInput       = "invalid input"
	InvalidEmail       = "invalid email"
	EmailAlreadyExists = "email already exists"
	EmptyPassword      = "password can't be empty"
	InvalidAuthToken   = "invalid auth-token"
	InvalidCredentials = "invalid credentials"
	Unauthorized       = "Unauthorized"
)

var (
	ErrInvalidInput = errors.New(InvalidInput)
	ErrInvalidEmail       = errors.New(InvalidEmail)
	ErrEmailAlreadyExists = errors.New(EmailAlreadyExists)
	ErrEmptyPassword      = errors.New(EmptyPassword)
	ErrInvalidAuthToken   = errors.New(InvalidAuthToken)
	ErrInvalidCredentials = errors.New(InvalidCredentials)
	ErrUnauthorized       = errors.New(Unauthorized)
)

var errorCode = map[string]int{
	InvalidInput: 		1,
	InvalidEmail:       2,
	EmailAlreadyExists: 3,
	EmptyPassword:      4,
	InvalidAuthToken:   5,
	InvalidCredentials: 6,
	Unauthorized:       7,
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
