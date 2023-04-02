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
	// Try to pull this form from the database.
	dbf, err := s.s.Form.GetByID(id)
	if err != nil {
		return nil, err
	}

	// Convert interface to modules.
	m, err := s.InterfaceToModules(dbf.Modules.([]interface{}))
	if err != nil {
		return nil, err
	}

	// Create a new Form.
	f := &types.Form{
		ID:      dbf.ID,
		Modules: m,
	}

	return f, nil
}

// InterfaceToModules takes in an []interface{} and converts it to individual
// module types.
func (s *Service) InterfaceToModules(i []interface{}) ([]types.Module, error) {
	// Create modules slice.
	modules := []types.Module{}

	// Loop through the interface slice.
	for _, v := range i {
		// Convert to map[string]interface{}.
		m := v.(map[string]interface{})

		// Get type.
		t, ok := m["type"]
		if !ok {
			return nil, errors.New("missing type")
		}

		// Switch type.
		switch t {
		case "short-text":
			// TODO: Should probably fully map this, checking if
			//       each field exists in the map and setting it
			//       directly.
			module := &types.ShortText{}

			// Handle type.
			typeStr, ok := t.(string)
			if !ok {
				return nil, errors.New("invalid type property, must be string")
			}
			module.Type = typeStr

			// Handle properties.
			p, ok := m["properties"]
			if !ok {
				return nil, errors.New("missing properties")
			}
			pm := p.(map[string]interface{})

			properties := types.ShortTextProperties{}

			// Handle property label.
			label, ok := pm["label"]
			if !ok {
				return nil, errors.New("missing property label")
			}
			labelStr, ok := label.(string)
			if !ok {
				return nil, errors.New("invalid property label, must be string")
			}
			properties.Label = labelStr

			// Set the properties.
			module.Properties = properties

			modules = append(modules, module)
		default:
			return nil, errors.New("invalid module type")
		}
	}

	return modules, nil
}
