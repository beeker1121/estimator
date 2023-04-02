package form

import "errors"

var (
	// ErrFormNotFound is returned when a form could not be found.
	ErrFormNotFound = errors.New("form could not be found")
)
