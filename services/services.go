package services

import (
	"estimator/services/form"
	"estimator/services/users"
	"estimator/storage"
)

// Services defines the main business logic services.
type Services struct {
	Users *users.Service
	Form  *form.Service
}

// New creates a new services.
func New(s *storage.Storage) *Services {
	return &Services{
		Users: users.New(s),
		Form:  form.New(s),
	}
}
