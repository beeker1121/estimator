package storage

import (
	"estimator/storage/form"
	"estimator/storage/users"
)

// Storage defines the storage system.
type Storage struct {
	Users users.Database
	Form  form.Database
}

// New returns a new storage.
func New(s *Storage) *Storage {
	return s
}
