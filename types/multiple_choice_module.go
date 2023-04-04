package types

// MultipleChoice defines the multiple choice module.
type MultipleChoice struct {
	ID         string                   `json:"id"`
	Type       string                   `json:"type"`
	Name       string                   `json:"name"`
	Properties MultipleChoiceProperties `json:"properties"`
}

// GetType implements the Module interface.
func (mc *MultipleChoice) GetType() string {
	return mc.Type
}

// Validate implements the Module interface.
func (mc *MultipleChoice) Validate() error {
	// Check type.
	if err := ValidateType(mc.Type); err != nil {
		return err
	}

	return nil
}

// MultipleChoiceProperties defines the multiple choice module properties.
type MultipleChoiceProperties struct {
	Label       string                 `json:"label"`
	Sublabel    string                 `json:"sublabel"`
	Tooltip     string                 `json:"tooltip"`
	Required    bool                   `json:"required"`
	Placeholder string                 `json:"placeholder"`
	Suffix      string                 `json:"suffix"`
	WidthType   bool                   `json:"width_type"`
	Width       int                    `json:"width"`
	Options     []MultipleChoiceOption `json:"options"`
}

// MultipleChoiceOption defines a multiple choice option.
type MultipleChoiceOption struct {
	ID    string `json:"id"`
	Value string `json:"value"`
}
