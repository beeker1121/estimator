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

	// Set ID and password.
	u.ID = uuid.NewString()
	u.Password = string(pwHash)

	// Map to storage type.
	su := &users.User{
		ID:        u.ID,
		AccountID: u.AccountID,
		Name:      u.Name,
		Email:     u.Email,
		Password:  u.Password,
		Role:      u.Role,
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
		ID:        su.ID,
		AccountID: su.AccountID,
		Name:      su.Name,
		Email:     su.Email,
		Password:  su.Password,
		Role:      su.Role,
	}

	return u, nil
}

// GetByID gets a user by the given ID.
func (s *Service) GetByID(id string) (*types.User, error) {
	// Try to pull this user from the database.
	su, err := s.s.Users.GetByID(id)
	if err != nil {
		return nil, err
	}

	// Map to user type.
	u := &types.User{
		ID:        su.ID,
		AccountID: su.AccountID,
		Name:      su.Name,
		Email:     su.Email,
		Password:  su.Password,
		Role:      su.Role,
	}

	return u, nil
}

// UpdateByID updates a user by the given ID.
func (s *Service) UpdateByID(id string, uup *types.UserUpdateParams) (*types.User, error) {
	// Create a new ParamErrors.
	pes := errors.NewParamErrors()

	// Map to storage type.
	sup := &users.UpdateParams{
		ID:        uup.ID,
		AccountID: uup.AccountID,
		Name:      uup.Name,
		Email:     uup.Email,
		Role:      uup.Role,
	}

	// Check email.
	if uup.Email != nil {
		if *uup.Email == "" {
			pes.Add(errors.NewParamError("email", ErrEmailEmpty))
		} else {
			_, err := s.s.Users.GetByEmail(*uup.Email)
			if err == nil {
				pes.Add(errors.NewParamError("email", ErrEmailExists))
			} else if err != nil && err != users.ErrUserNotFound {
				return nil, err
			}
		}
	}

	// Check password.
	if uup.Password != nil {
		if len(*uup.Password) < 8 {
			pes.Add(errors.NewParamError("password", ErrPassword))
		}

		// Return if there were parameter errors.
		if pes.Length() > 0 {
			return nil, pes
		}

		// Hash the password.
		pwHash, err := bcrypt.GenerateFromPassword([]byte(*uup.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}
		pwHashStr := string(pwHash)

		sup.Password = &pwHashStr
	}

	// Update in storage.
	su, err := s.s.Users.UpdateByID(id, sup)
	if err != nil {
		return nil, err
	}

	// Map to user type.
	u := &types.User{
		ID:        su.ID,
		AccountID: su.AccountID,
		Name:      su.Name,
		Email:     su.Email,
		Password:  su.Password,
		Role:      su.Role,
	}

	return u, nil
}
