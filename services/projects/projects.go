package projects

import (
	"estimator/services/errors"
	"estimator/storage"
	"estimator/storage/projects"
	"estimator/types"

	"github.com/google/uuid"
)

// Service defines the projects service.
type Service struct {
	s *storage.Storage
}

// New creates a new service.
func New(s *storage.Storage) *Service {
	return &Service{
		s: s,
	}
}

// Create creates a new project.
func (s *Service) Create(p *types.Project) (*types.Project, error) {
	// Create a new ParamErrors.
	pes := errors.NewParamErrors()

	// Check account ID.
	//
	// TODO: Move this to a Validate() function.
	if p.AccountID == "" {
		pes.Add(errors.NewParamError("account_id", ErrAccountIDEmpty))
	}

	// Check name.
	//
	// TODO: Move this to a Validate() function.
	if p.Name == "" {
		pes.Add(errors.NewParamError("name", ErrNameEmpty))
	}

	// Return if there were parameter errors.
	if pes.Length() > 0 {
		return nil, pes
	}

	// Set ID.
	p.ID = uuid.NewString()

	// Map to storage type.
	sp := &projects.Project{
		ID:        p.ID,
		AccountID: p.AccountID,
		Name:      p.Name,
	}

	// Create in storage.
	var err error
	sp, err = s.s.Projects.Create(sp)
	if err != nil {
		return nil, err
	}

	return p, nil
}

// GetByID gets a project by the given ID.
func (s *Service) GetByID(id string) (*types.Project, error) {
	// Try to pull this project from the database.
	sp, err := s.s.Projects.GetByID(id)
	if err != nil {
		return nil, err
	}

	// Map to project type.
	p := &types.Project{
		ID:        sp.ID,
		AccountID: sp.AccountID,
		Name:      sp.Name,
	}

	return p, nil
}

// GetByIDAndUserID gets a project by the given ID and user ID.
func (s *Service) GetByIDAndUserID(id, userID string) (*types.Project, error) {
	// Get this user from storage to verify.
	//
	// TODO: Implement this.
	// _, err := s.s.Users.GetByIDAndAccountID(userID, id)
	// if err == users.ErrUserNotFound {
	// 	return nil, ErrAccountNotFound
	// } else if err != nil {
	// 	return nil, err
	// }

	// TODO: Verify role.

	// Try to pull this project from the database.
	sp, err := s.s.Projects.GetByID(id)
	if err != nil {
		return nil, err
	}

	// Map to project type.
	p := &types.Project{
		ID:        sp.ID,
		AccountID: sp.AccountID,
		Name:      sp.Name,
	}

	return p, nil
}

// UpdateByID updates a project by the given ID.
func (s *Service) UpdateByID(id string, pup *types.ProjectUpdateParams) (*types.Project, error) {
	// Create a new ParamErrors.
	pes := errors.NewParamErrors()

	// Check account ID.
	//
	// TODO: Move this to a Validate() function.
	if pup.AccountID != nil {
		if *pup.AccountID == "" {
			pes.Add(errors.NewParamError("account_id", ErrAccountIDEmpty))
		}
	}

	// Check name.
	//
	// TODO: Move this to a Validate() function.
	if pup.Name != nil {
		if *pup.Name == "" {
			pes.Add(errors.NewParamError("name", ErrNameEmpty))
		}
	}

	// Return if there were parameter errors.
	if pes.Length() > 0 {
		return nil, pes
	}

	// Map to storage type.
	sup := &projects.UpdateParams{
		AccountID: pup.AccountID,
		Name:      pup.Name,
	}

	// Update in storage.
	sp, err := s.s.Projects.UpdateByID(id, sup)
	if err != nil {
		return nil, err
	}

	// Map to project type.
	p := &types.Project{
		ID:        sp.ID,
		AccountID: sp.AccountID,
		Name:      sp.Name,
	}

	return p, nil
}

// UpdateByIDAndUserID updates a project by the given ID and user ID.
func (s *Service) UpdateByIDAndUserID(id, userID string, pup *types.ProjectUpdateParams) (*types.Project, error) {
	// Create a new ParamErrors.
	pes := errors.NewParamErrors()

	// Check account ID.
	//
	// TODO: Move this to a Validate() function.
	if pup.AccountID != nil {
		if *pup.AccountID == "" {
			pes.Add(errors.NewParamError("account_id", ErrAccountIDEmpty))
		}
	}

	// Check name.
	//
	// TODO: Move this to a Validate() function.
	if pup.Name != nil {
		if *pup.Name == "" {
			pes.Add(errors.NewParamError("name", ErrNameEmpty))
		}
	}

	// Return if there were parameter errors.
	if pes.Length() > 0 {
		return nil, pes
	}

	// Get this user from storage to verify.
	//
	// TODO: Implement this.
	// _, err := s.s.Users.GetByIDAndAccountID(userID, id)
	// if err == users.ErrUserNotFound {
	// 	return nil, ErrAccountNotFound
	// } else if err != nil {
	// 	return nil, err
	// }

	// TODO: Verify role.

	// Map to storage type.
	sup := &projects.UpdateParams{
		AccountID: pup.AccountID,
		Name:      pup.Name,
	}

	// Update in storage.
	sp, err := s.s.Projects.UpdateByID(id, sup)
	if err != nil {
		return nil, err
	}

	// Map to project type.
	p := &types.Project{
		ID:        sp.ID,
		AccountID: sp.AccountID,
		Name:      sp.Name,
	}

	return p, nil
}
