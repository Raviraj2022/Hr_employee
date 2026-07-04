package errors

import "errors"

var (
	ErrEmailExists = errors.New("email already exists")
	ErrInvalidLogin = errors.New("invalid email or password")
	ErrUserNotFound = errors.New("user not found")
)