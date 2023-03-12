package types

// ShortText defines the short text module.
type ShortText struct {
	ID         string              `json:"id"`
	Type       string              `json:"type"`
	Name       string              `json:"name"`
	Properties ShortTextProperties `json:"properties"`
}

// GetType implements the Module interface.
func (st *ShortText) GetType() string {
	return st.Type
}

// Validate implements the Module interface.
func (st *ShortText) Validate() error {
	// Check type.
	if err := ValidateType(st.Type); err != nil {
		return err
	}

	return nil
}

// ShortTextProperties defines the short text module properties.
type ShortTextProperties struct {
	Label       string `json:"label"`
	Sublabel    string `json:"sublabel"`
	Tooltip     string `json:"tooltip"`
	Required    bool   `json:"required"`
	Placeholder string `json:"placeholder"`
	Suffix      string `json:"suffix"`
	WidthType   bool   `json:"width_type"`
	Width       int    `json:"width"`
	Validation  string `json:"validation"`
}
