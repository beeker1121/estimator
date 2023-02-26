package form

import (
	"estimator/storage"
	"estimator/types"
)

// Service defines the form service.
type Service struct {
	s *storage.Storage
}

// New creates a new service.
func New(s *storage.Storage) *Service {
	return &Service{
		s: s,
	}
}

// Create creates a new form.
func (s *Service) Create(f *types.Form) (*types.Form, error) {
	return &types.Form{}, nil
}

// GetByID gets a form by the given ID.
func (s *Service) GetByID(id string) (*types.Form, error) {
	return &types.Form{}, nil
}
