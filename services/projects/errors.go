package projects

import "errors"

var (
	// ErrAccountIDEmpty is returned when the account ID is empty.
	ErrAccountIDEmpty = errors.New("account ID is empty")

	// ErrNameEmpty is returned when the name is empty.
	ErrNameEmpty = errors.New("name is empty")

	// ErrProjectNotFound is returned when a project could not be found.
	ErrProjectNotFound = errors.New("project could not be found")
)
