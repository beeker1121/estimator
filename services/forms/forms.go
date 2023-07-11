package forms

import (
	"errors"

	"estimator/storage"
	"estimator/storage/forms"
	"estimator/types"

	"github.com/google/uuid"
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

	// Set ID.
	f.ID = uuid.NewString()

	// Loop through the modules.
	for i, module := range f.Modules {
		// Validate the module.
		if err := module.Validate(); err != nil {
			return nil, err
		}

		// Set ID.
		f.Modules[i].SetID(uuid.NewString())
	}

	// TODO: Validate project ID.

	// TODO: Validate name.

	// TODO: Validate properties.

	// TODO: Validate button.

	// Map to storage type.
	sf := &forms.Form{
		ID:         f.ID,
		ProjectID:  f.ProjectID,
		Name:       f.Name,
		Properties: f.Properties,
		Button:     f.Button,
		Modules:    f.Modules,
	}

	// Create in storage.
	sf, err = s.s.Forms.Create(sf)
	if err != nil {
		return nil, err
	}

	return f, nil
}

// GetByID gets a form by the given ID.
func (s *Service) GetByID(id string) (*types.Form, error) {
	// Try to pull this form from the database.
	sf, err := s.s.Forms.GetByID(id)
	if err == forms.ErrFormNotFound {
		return nil, ErrFormNotFound
	} else if err != nil {
		return nil, err
	}

	// Convert properties.
	propertiesMap := sf.Properties.(map[string]interface{})
	properties := types.FormProperties{
		BackgroundColor: propertiesMap["background_color"].(string),
		FontColor:       propertiesMap["font_color"].(string),
	}

	// Convert button.
	buttonMap := sf.Button.(map[string]interface{})
	button := types.FormButton{
		BackgroundColor: buttonMap["background_color"].(string),
		Color:           buttonMap["color"].(string),
		FontSize:        buttonMap["font_size"].(string),
		FontFamily:      buttonMap["font_family"].(string),
	}

	// Convert interface to modules.
	m, err := s.InterfaceToModules(sf.Modules.([]interface{}))
	if err != nil {
		return nil, err
	}

	// Create a new Form.
	f := &types.Form{
		ID:         sf.ID,
		ProjectID:  sf.ProjectID,
		Name:       sf.Name,
		Properties: properties,
		Button:     button,
		Modules:    m,
	}

	return f, nil
}

// UpdateByIDAndUserID updates a form by the given ID and member ID.
func (s *Service) UpdateByIDAndUserID(id, userID string, f *types.Form) (*types.Form, error) {
	var err error

	// Loop through the modules.
	for _, module := range f.Modules {
		// Validate the module.
		if err := module.Validate(); err != nil {
			return nil, err
		}
	}

	// TODO: Validate project ID.

	// TODO: Validate name.

	// TODO: Validate properties.

	// TODO: Validate button.

	// Map to storage type.
	sf := &forms.Form{
		ID:         id,
		ProjectID:  f.ProjectID,
		Name:       f.Name,
		Properties: f.Properties,
		Button:     f.Button,
		Modules:    f.Modules,
	}

	// Update in storage.
	sf, err = s.s.Forms.UpdateByID(id, sf)
	if err != nil {
		return nil, err
	}

	return f, nil
}

// InterfaceToModules takes in an []interface{} and converts it to individual
// module types.
//
// TODO: This should return a services.ParamErrors type on validation errors.
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
		case "heading":
			// TODO: Should probably fully map this, checking if
			//       each field exists in the map and setting it
			//       directly.
			module := &types.Heading{}

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

			properties := types.HeadingProperties{}

			// Handle property title.
			title, ok := pm["title"]
			if !ok {
				return nil, errors.New("missing property title")
			}
			titleStr, ok := title.(string)
			if !ok {
				return nil, errors.New("invalid property title, must be a string")
			}
			properties.Title = titleStr

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

			// Handle property size.
			size, ok := pm["size"]
			if !ok {
				return nil, errors.New("missing property size")
			}
			sizeStr, ok := size.(string)
			if !ok {
				return nil, errors.New("invalid property size, must be a string")
			}
			properties.Size = sizeStr

			// Handle property alignment.
			alignment, ok := pm["alignment"]
			if !ok {
				return nil, errors.New("missing property alignment")
			}
			alignmentStr, ok := alignment.(string)
			if !ok {
				return nil, errors.New("invalid property alignment, must be a string")
			}
			properties.Alignment = alignmentStr

			// Handle property imageAlignment.
			imageAlignment, ok := pm["image_alignment"]
			if !ok {
				return nil, errors.New("missing property image alignment")
			}
			imageAlignmentStr, ok := imageAlignment.(string)
			if !ok {
				return nil, errors.New("invalid property image alignment, must be string")
			}
			properties.ImageAlignment = imageAlignmentStr

			// Handle property verticalAlignment.
			verticalAlignment, ok := pm["vertical_alignment"]
			if !ok {
				return nil, errors.New("missing property vertical alignment")
			}
			verticalAlignmentStr, ok := verticalAlignment.(string)
			if !ok {
				return nil, errors.New("invalid property vertical alignment, must be string")
			}
			properties.VerticalAlignment = verticalAlignmentStr

			// TODO: Handle property image.

			// Handle property imageWidth.
			imageWidth, ok := pm["image_width"]
			if !ok {
				return nil, errors.New("missing property image width")
			}
			imageWidthFloat64, ok := imageWidth.(float64)
			if !ok {
				return nil, errors.New("invalid property image width, must be an integer")
			}
			properties.ImageWidth = int(imageWidthFloat64)

			// Set the properties.
			module.Properties = properties

			modules = append(modules, module)
		case "full-name":
			// TODO: Should probably fully map this, checking if
			//       each field exists in the map and setting it
			//       directly.
			module := &types.FullName{}

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

			properties := types.FullNameProperties{}

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

			// Handle property show prefix.
			showPrefix, ok := pm["show_prefix"]
			if !ok {
				return nil, errors.New("missing property show prefix")
			}
			showPrefixBool, ok := showPrefix.(bool)
			if !ok {
				return nil, errors.New("invalid property show prefix, must be a boolean")
			}
			properties.ShowPrefix = showPrefixBool

			// Handle property show middle name.
			showMiddleName, ok := pm["show_middle_name"]
			if !ok {
				return nil, errors.New("missing property show middle name")
			}
			showMiddleNameBool, ok := showMiddleName.(bool)
			if !ok {
				return nil, errors.New("invalid property show middle name, must be a boolean")
			}
			properties.ShowMiddleName = showMiddleNameBool

			// Handle property show suffix.
			showSuffix, ok := pm["show_suffix"]
			if !ok {
				return nil, errors.New("missing property show suffix")
			}
			showSuffixBool, ok := showSuffix.(bool)
			if !ok {
				return nil, errors.New("invalid property show suffix, must be a boolean")
			}
			properties.ShowSuffix = showSuffixBool

			// Handle property prefix sublabel.
			prefixSublabel, ok := pm["prefix_sublabel"]
			if !ok {
				return nil, errors.New("missing property prefix sublabel")
			}
			prefixSublabelStr, ok := prefixSublabel.(string)
			if !ok {
				return nil, errors.New("invalid property prefix sublabel, must be a string")
			}
			properties.PrefixSublabel = prefixSublabelStr

			// Handle property first name sublabel.
			firstNameSublabel, ok := pm["first_name_sublabel"]
			if !ok {
				return nil, errors.New("missing property first name sublabel")
			}
			firstNameSublabelStr, ok := firstNameSublabel.(string)
			if !ok {
				return nil, errors.New("invalid property first name sublabel, must be a string")
			}
			properties.FirstNameSublabel = firstNameSublabelStr

			// Handle property middle name sublabel.
			middleNameSublabel, ok := pm["middle_name_sublabel"]
			if !ok {
				return nil, errors.New("missing property middle name sublabel")
			}
			middleNameSublabelStr, ok := middleNameSublabel.(string)
			if !ok {
				return nil, errors.New("invalid property middle name sublabel, must be a string")
			}
			properties.MiddleNameSublabel = middleNameSublabelStr

			// Handle property last name sublabel.
			lastNameSublabel, ok := pm["last_name_sublabel"]
			if !ok {
				return nil, errors.New("missing property last name sublabel")
			}
			lastNameSublabelStr, ok := lastNameSublabel.(string)
			if !ok {
				return nil, errors.New("invalid property last name sublabel, must be a string")
			}
			properties.LastNameSublabel = lastNameSublabelStr

			// Handle property suffix sublabel.
			suffixSublabel, ok := pm["suffix_sublabel"]
			if !ok {
				return nil, errors.New("missing property suffix sublabel")
			}
			suffixSublabelStr, ok := suffixSublabel.(string)
			if !ok {
				return nil, errors.New("invalid property suffix sublabel, must be a string")
			}
			properties.SuffixSublabel = suffixSublabelStr

			// Set the properties.
			module.Properties = properties

			modules = append(modules, module)
		default:
			return nil, errors.New("invalid module type")
		}
	}

	return modules, nil
}
