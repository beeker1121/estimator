package form

import (
	"errors"
	"estimator/storage"
	"estimator/storage/form"
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
	var err error

	// Loop through the modules.
	for _, module := range f.Modules {
		// Validate the module.
		if err := module.Validate(); err != nil {
			return nil, err
		}
	}

	// Map to storage type.
	sf := &form.Form{
		Modules: f.Modules,
	}

	// Create in storage.
	sf, err = s.s.Form.Create(sf)
	if err != nil {
		return nil, err
	}

	// Update type.
	f.ID = sf.ID

	return f, nil
}

// GetByID gets a form by the given ID.
func (s *Service) GetByID(id string) (*types.Form, error) {
	return &types.Form{}, nil
}

// InterfaceToModules takes in a map[string]interface{} and converts it to
// individual modules types.
func (s *Service) InterfaceToModules(i []interface{}) ([]types.Module, error) {
	// Create modules slice.
	modules := []types.Module{}

	// Loop through the interface slice.
	for _, v := range i {
		// Convert to map[string]interface{}.
		m := v.(map[string]interface{})

		// Switch type.
		switch m["type"] {
		case "short-text":
			module := &types.ShortText{
				Type: m["type"].(string),
			}

			modules = append(modules, module)
		default:
			return nil, errors.New("invalid module type")
		}
	}

	return modules, nil
}
