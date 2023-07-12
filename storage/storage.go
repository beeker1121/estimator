package storage

import (
	"estimator/storage/accounts"
	"estimator/storage/forms"
	"estimator/storage/projects"
	"estimator/storage/users"
)

// Storage defines the storage system.
type Storage struct {
	Accounts accounts.Database
	Users    users.Database
	Projects projects.Database
	Forms    forms.Database
}

// New returns a new storage.
func New(s *Storage) *Storage {
	return s
}
