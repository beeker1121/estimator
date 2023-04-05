package types

// FullName defines the full name module.
type FullName struct {
	ID         string             `json:"id"`
	Type       string             `json:"type"`
	Name       string             `json:"name"`
	Properties FullNameProperties `json:"properties"`
}

// SetID implements the Module interface.
func (fn *FullName) SetID(id string) {
	fn.ID = id
}

// GetType implements the Module interface.
func (fn *FullName) GetType() string {
	return fn.Type
}

// Validate implements the Module interface.
func (fn *FullName) Validate() error {
	// Check type.
	if err := ValidateType(fn.Type); err != nil {
		return err
	}

	return nil
}

// FullNameProperties defines the full name module properties.
type FullNameProperties struct {
	Label              string `json:"label"`
	Tooltip            string `json:"tooltip"`
	Required           bool   `json:"required"`
	ShowPrefix         bool   `json:"show_prefix"`
	ShowMiddleName     bool   `json:"show_middle_name"`
	ShowSuffix         bool   `json:"show_suffix"`
	PrefixSublabel     string `json:"prefix_sublabel"`
	FirstNameSublabel  string `json:"first_name_sublabel"`
	MiddleNameSublabel string `json:"middle_name_sublabel"`
	LastNameSublabel   string `json:"last_name_sublabel"`
	SuffixSublabel     string `json:"suffix_sublabel"`
}
