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

// UpdateByIDAndMemberID updates a form by the given ID and member ID.
func (s *Service) UpdateByIDAndMemberID(id, memberID string, f *types.Form) (*types.Form, error) {
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
	sf, err = s.s.Form.UpdateByID(id, sf)
	if err != nil {
		return nil, err
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
				return nil, errors.New("invalid type property, must be a string")
			}
			module.Type = typeStr

			// Handle name.
			name, ok := m["name"]
			if !ok {
				return nil, errors.New("missing name for module")
			}
			nameStr, ok := name.(string)
			if !ok {
				return nil, errors.New("invalid module name, must be a string")
			}
			module.Name = nameStr

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
				return nil, errors.New("invalid property label, must be a string")
			}
			properties.Label = labelStr

			// Handle property sublabel.
			sublabel, ok := pm["sublabel"]
			if !ok {
				return nil, errors.New("missing property sublabel")
			}
			sublabelStr, ok := sublabel.(string)
			if !ok {
				return nil, errors.New("invalid property sublabel, must be a string")
			}
			properties.Sublabel = sublabelStr

			// Handle property tooltip.
			tooltip, ok := pm["tooltip"]
			if !ok {
				return nil, errors.New("missing property tooltip")
			}
			tooltipStr, ok := tooltip.(string)
			if !ok {
				return nil, errors.New("invalid property tooltip, must be a string")
			}
			properties.Tooltip = tooltipStr

			// Handle property required.
			required, ok := pm["required"]
			if !ok {
				return nil, errors.New("missing property required")
			}
			requiredBool, ok := required.(bool)
			if !ok {
				return nil, errors.New("invalid property required, must be a boolean")
			}
			properties.Required = requiredBool

			// Handle property placeholder.
			placeholder, ok := pm["placeholder"]
			if !ok {
				return nil, errors.New("missing property placeholder")
			}
			placeholderStr, ok := placeholder.(string)
			if !ok {
				return nil, errors.New("invalid property placeholder, must be string")
			}
			properties.Placeholder = placeholderStr

			// Handle property default.
			defaultp, ok := pm["default"]
			if !ok {
				return nil, errors.New("missing property default")
			}
			defaultpStr, ok := defaultp.(string)
			if !ok {
				return nil, errors.New("invalid property default, must be string")
			}
			properties.Default = defaultpStr

			// Handle property suffix.
			suffix, ok := pm["suffix"]
			if !ok {
				return nil, errors.New("missing property suffix")
			}
			suffixStr, ok := suffix.(string)
			if !ok {
				return nil, errors.New("invalid property suffix, must be a string")
			}
			properties.Suffix = suffixStr

			// Handle property widthType.
			widthType, ok := pm["width_type"]
			if !ok {
				return nil, errors.New("missing property width type")
			}
			widthTypeBool, ok := widthType.(bool)
			if !ok {
				return nil, errors.New("invalid property width type, must be a boolean")
			}
			properties.WidthType = widthTypeBool

			// Handle property width.
			width, ok := pm["width"]
			if !ok {
				return nil, errors.New("missing property width")
			}
			widthFloat64, ok := width.(float64)
			if !ok {
				return nil, errors.New("invalid property width, must be an integer")
			}
			properties.Width = int(widthFloat64)

			// Handle property validation.
			validation, ok := pm["validation"]
			if !ok {
				return nil, errors.New("missing property validation")
			}
			validationStr, ok := validation.(string)
			if !ok {
				return nil, errors.New("invalid property validation type, must be string")
			}
			properties.Validation = validationStr

			// Set the properties.
			module.Properties = properties

			modules = append(modules, module)
		case "multiple-choice":
			// TODO: Should probably fully map this, checking if
			//       each field exists in the map and setting it
			//       directly.
			module := &types.MultipleChoice{}

			// Handle type.
			typeStr, ok := t.(string)
			if !ok {
				return nil, errors.New("invalid type property, must be a string")
			}
			module.Type = typeStr

			// Handle name.
			name, ok := m["name"]
			if !ok {
				return nil, errors.New("missing name for module")
			}
			nameStr, ok := name.(string)
			if !ok {
				return nil, errors.New("invalid module name, must be a string")
			}
			module.Name = nameStr

			// Handle properties.
			p, ok := m["properties"]
			if !ok {
				return nil, errors.New("missing properties")
			}
			pm := p.(map[string]interface{})

			properties := types.MultipleChoiceProperties{}

			// Handle property label.
			label, ok := pm["label"]
			if !ok {
				return nil, errors.New("missing property label")
			}
			labelStr, ok := label.(string)
			if !ok {
				return nil, errors.New("invalid property label, must be a string")
			}
			properties.Label = labelStr

			// Handle property sublabel.
			sublabel, ok := pm["sublabel"]
			if !ok {
				return nil, errors.New("missing property sublabel")
			}
			sublabelStr, ok := sublabel.(string)
			if !ok {
				return nil, errors.New("invalid property sublabel, must be a string")
			}
			properties.Sublabel = sublabelStr

			// Handle property tooltip.
			tooltip, ok := pm["tooltip"]
			if !ok {
				return nil, errors.New("missing property tooltip")
			}
			tooltipStr, ok := tooltip.(string)
			if !ok {
				return nil, errors.New("invalid property tooltip, must be a string")
			}
			properties.Tooltip = tooltipStr

			// Handle property required.
			required, ok := pm["required"]
			if !ok {
				return nil, errors.New("missing property required")
			}
			requiredBool, ok := required.(bool)
			if !ok {
				return nil, errors.New("invalid property required, must be a boolean")
			}
			properties.Required = requiredBool

			// Handle property placeholder.
			placeholder, ok := pm["placeholder"]
			if !ok {
				return nil, errors.New("missing property placeholder")
			}
			placeholderStr, ok := placeholder.(string)
			if !ok {
				return nil, errors.New("invalid property placeholder, must be string")
			}
			properties.Placeholder = placeholderStr

			// Handle property suffix.
			suffix, ok := pm["suffix"]
			if !ok {
				return nil, errors.New("missing property suffix")
			}
			suffixStr, ok := suffix.(string)
			if !ok {
				return nil, errors.New("invalid property suffix, must be a string")
			}
			properties.Suffix = suffixStr

			// Handle property widthType.
			widthType, ok := pm["width_type"]
			if !ok {
				return nil, errors.New("missing property width_type")
			}
			widthTypeBool, ok := widthType.(bool)
			if !ok {
				return nil, errors.New("invalid property width type, must be a boolean")
			}
			properties.WidthType = widthTypeBool

			// Handle property width.
			width, ok := pm["width"]
			if !ok {
				return nil, errors.New("missing property width")
			}
			widthFloat64, ok := width.(float64)
			if !ok {
				return nil, errors.New("invalid property width, must be an integer")
			}
			properties.Width = int(widthFloat64)

			// Handle property options.
			options, ok := pm["options"]
			if !ok {
				return nil, errors.New("missing property options")
			}
			optionsSlice, ok := options.([]interface{})
			if !ok {
				return nil, errors.New("multiple choice options is not an array of objects")
			}

			var optionsType []types.MultipleChoiceOption
			for _, v := range optionsSlice {
				option, ok := v.(map[string]interface{})
				if !ok {
					return nil, errors.New("option is invalid")
				}

				optionID, ok := option["id"]
				if !ok {
					return nil, errors.New("could not get 'id' property of multiple choice option")
				}
				optionIDStr, ok := optionID.(string)
				if !ok {
					return nil, errors.New("invalid option ID, must be a string")
				}

				optionValue, ok := option["id"]
				if !ok {
					return nil, errors.New("could not get 'value' property of multiple choice option")
				}
				optionValueStr, ok := optionValue.(string)
				if !ok {
					return nil, errors.New("invalid option value, must be a string")
				}

				optionsType = append(optionsType, types.MultipleChoiceOption{
					ID:    optionIDStr,
					Value: optionValueStr,
				})

			}
			properties.Options = optionsType

			// Set the properties.
			module.Properties = properties

			modules = append(modules, module)
		default:
			return nil, errors.New("invalid module type")
		}
	}

	return modules, nil
}
