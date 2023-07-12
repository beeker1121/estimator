package services

import (
	"estimator/services/accounts"
	"estimator/services/forms"
	"estimator/services/projects"
	"estimator/services/users"
	"estimator/storage"
)

// Services defines the main business logic services.
type Services struct {
	Accounts *accounts.Service
	Users    *users.Service
	Projects *projects.Service
	Forms    *forms.Service
}

// New creates a new services.
func New(s *storage.Storage) *Services {
	return &Services{
		Accounts: accounts.New(s),
		Users:    users.New(s),
		Projects: projects.New(s),
		Forms:    forms.New(s),
	}
}
