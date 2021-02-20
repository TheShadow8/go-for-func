package util

type AppError struct {
	error   error
	status  int
	success bool
}

func NewAppError(error error, status int) error {
	success := true

	if error != nil {

		success = false
	}

	e := AppError{error, status, success}
	return &e
}

func (ae *AppError) Error() string {
	return ae.error.Error()
}

func (ae *AppError) GetStatus() int {
	return ae.status
}
