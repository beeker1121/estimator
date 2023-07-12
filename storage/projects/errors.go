package projects

import "errors"

var (
	// ErrProjectNotFound is returned when an project could not be found.
	ErrProjectNotFound = errors.New("project could not be found")
)
