package accounts

import "errors"

var (
	// ErrNameEmpty is returned when the name is empty.
	ErrNameEmpty = errors.New("name is empty")

	// ErrAccountNotFound is returned when an account could not be found.
	ErrAccountNotFound = errors.New("account could not be found")
)
