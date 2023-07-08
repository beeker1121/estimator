package accounts

import "errors"

var (
	// ErrAccountNotFound is returned when an account could not be found.
	ErrAccountNotFound = errors.New("account could not be found")
)
