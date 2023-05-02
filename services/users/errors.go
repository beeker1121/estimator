package users

import "errors"

var (
	// ErrEmailEmpty is returned when the email is empty.
	ErrEmailEmpty = errors.New("email is empty")

	// ErrEmailExists is returned when the email already exists.
	ErrEmailExists = errors.New("email already exists")

	// ErrPassword is returned when the password is in an invalid format.
	ErrPassword = errors.New("password must be at least 8 characters")

	// ErrInvalidLogin is returned when the email and/or password used
	// with login is invalid.
	ErrInvalidLogin = errors.New("email and/or password is invalid")

	// ErrUserNotFound is returned when a user could not be found.
	ErrUserNotFound = errors.New("user could not be found")
)