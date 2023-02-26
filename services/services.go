package services

import (
	"estimator/services/form"
	"estimator/storage"
)

// Services defines the main business logic services.
type Services struct {
	Form *form.Service
}

// New creates a new services.
func New(s *storage.Storage) *Services {
	return &Services{
		Form: form.New(s),
	}
}
