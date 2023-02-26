package storage

import (
	"estimator/storage/form"
)

// Storage defines the storage system.
type Storage struct {
	Form form.Database
}

// New returns a new storage.
func New(s *Storage) *Storage {
	return s
}
