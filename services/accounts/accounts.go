package accounts

import (
	"estimator/services/errors"
	"estimator/storage"
	"estimator/storage/accounts"
	"estimator/types"

	"github.com/google/uuid"
)

// Service defines the accounts service.
type Service struct {
	s *storage.Storage
}

// New creates a new service.
func New(s *storage.Storage) *Service {
	return &Service{
		s: s,
	}
}

// Create creates a new account.
func (s *Service) Create(a *types.Account) (*types.Account, error) {
	// Create a new ParamErrors.
	pes := errors.NewParamErrors()

	// Check name.
	if a.Name == "" {
		pes.Add(errors.NewParamError("name", ErrNameEmpty))
	}

	// Set ID.
	a.ID = uuid.NewString()

	// Map to storage type.
	sa := &accounts.Account{
		ID:   a.ID,
		Name: a.Name,
	}

	// Create in storage.
	var err error
	sa, err = s.s.Accounts.Create(sa)
	if err != nil {
		return nil, err
	}

	return a, nil
}

// GetByID gets a account by the given ID.
func (s *Service) GetByID(id string) (*types.Account, error) {
	// Try to pull this account from the database.
	dba, err := s.s.Accounts.GetByID(id)
	if err != nil {
		return nil, err
	}

	// Map to account type.
	u := &types.Account{
		ID:   dba.ID,
		Name: dba.Name,
	}

	return u, nil
}

// UpdateByID updates a account by the given ID.
func (s *Service) UpdateByID(id string, a *types.Account) (*types.Account, error) {
	var err error

	// Map to storage type.
	sa := &accounts.Account{
		ID:   id,
		Name: a.Name,
	}

	// Create in storage.
	sa, err = s.s.Accounts.UpdateByID(id, sa)
	if err != nil {
		return nil, err
	}

	return a, nil
}

// UpdateByIDAndUserID updates a account by the given ID and user ID.
func (s *Service) UpdateByIDAndUserID(id, userID string, a *types.Account) (*types.Account, error) {
	var err error

	// Map to storage type.
	sa := &accounts.Account{
		ID:   id,
		Name: a.Name,
	}

	// Update in storage.
	sa, err = s.s.Accounts.UpdateByID(id, sa)
	if err != nil {
		return nil, err
	}

	return a, nil
}
