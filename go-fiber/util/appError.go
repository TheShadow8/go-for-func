package util

import "errors"

var (
	ErrInvalidEmail       = errors.New("invalid email")
	ErrEmailAlreadyExists = errors.New("email already exists")
	ErrEmptyPassword      = errors.New("password can't be empty")
	ErrInvalidAuthToken   = errors.New("invalid auth-token")
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrUnauthorized       = errors.New("Unauthorized")
)

type AppError struct {
	error error
	code  int
}

func NewAppError(error error, code int) error {
	e := AppError{error, code}
	return &e
}

func (ae *AppError) Error() string {
	return ae.error.Error()
}

func (ae *AppError) GetCode() int {
	return ae.code
}
