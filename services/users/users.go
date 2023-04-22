package users

import (
	"estimator/storage"
	"estimator/storage/users"
	"estimator/types"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// Service defines the user service.
type Service struct {
	s *storage.Storage
}

// New creates a new service.
func New(s *storage.Storage) *Service {
	return &Service{
		s: s,
	}
}

// Create creates a new user.
func (s *Service) Create(u *types.User) (*types.User, error) {
	var err error

	// Check email.
	if u.Email == "" {
		return nil, ErrEmailEmpty
	} else {
		_, err := s.s.Users.GetByEmail(u.Email)
		if err == nil {
			return nil, ErrEmailExists
		} else if err != nil && err != users.ErrUserNotFound {
			return nil, err
		}
	}

	// Check password.
	if len(u.Password) < 8 {
		return nil, ErrPassword
	}

	// Hash the password.
	pwHash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// Set ID.
	u.ID = uuid.NewString()

	// Map to storage type.
	su := &users.User{
		ID:       u.ID,
		Email:    u.Email,
		Password: string(pwHash),
	}

	// Create in storage.
	su, err = s.s.Users.Create(su)
	if err != nil {
		return nil, err
	}

	return u, nil
}

// GetByID gets a user by the given ID.
func (s *Service) GetByID(id string) (*types.User, error) {
	// Try to pull this user from the database.
	dbu, err := s.s.Users.GetByID(id)
	if err != nil {
		return nil, err
	}

	// Create a new User.
	u := &types.User{
		ID:       dbu.ID,
		Email:    dbu.Email,
		Password: dbu.Password,
	}

	return u, nil
}

// UpdateByID updates a user by the given ID.
func (s *Service) UpdateByID(id string, u *types.User) (*types.User, error) {
	var err error

	// Map to storage type.
	su := &users.User{
		ID:       id,
		Email:    u.Email,
		Password: u.Password,
	}

	// Create in storage.
	su, err = s.s.Users.UpdateByID(id, su)
	if err != nil {
		return nil, err
	}

	return u, nil
}
