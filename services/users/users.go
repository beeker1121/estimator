package users

import (
	"estimator/services/errors"
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
	// Create a new ParamErrors.
	pes := errors.NewParamErrors()

	// Check email.
	if u.Email == "" {
		pes.Add(errors.NewParamError("email", ErrEmailEmpty))
	} else {
		_, err := s.s.Users.GetByEmail(u.Email)
		if err == nil {
			pes.Add(errors.NewParamError("email", ErrEmailExists))
		} else if err != nil && err != users.ErrUserNotFound {
			return nil, err
		}
	}

	// Check password.
	if len(u.Password) < 8 {
		pes.Add(errors.NewParamError("password", ErrPassword))
	}

	// Return if there were parameter errors.
	if pes.Length() > 0 {
		return nil, pes
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

// Login checks if a user exists in the database and can log in.
func (s *Service) Login(u *types.User) (*types.User, error) {
	// Try to pull this user from the database.
	su, err := s.s.Users.GetByEmail(u.Email)
	if err == users.ErrUserNotFound {
		return nil, ErrInvalidLogin
	} else if err != nil {
		return nil, err
	}

	// Validate the password.
	if err = bcrypt.CompareHashAndPassword([]byte(su.Password), []byte(u.Password)); err != nil {
		return nil, ErrInvalidLogin
	}

	// Map to user type.
	u = &types.User{
		ID:       su.ID,
		Email:    su.Email,
		Password: su.Password,
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

	// Map to user type.
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
