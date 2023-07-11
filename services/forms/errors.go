package forms

import "errors"

var (
	// ErrNameEmpty is returned when the name is empty.
	ErrNameEmpty = errors.New("name is empty")

	// ErrFormNotFound is returned when a form could not be found.
	ErrFormNotFound = errors.New("form could not be found")
)
