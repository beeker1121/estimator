package types

// Heading defines the heading module.
type Heading struct {
	ID         string            `json:"id"`
	Type       string            `json:"type"`
	Name       string            `json:"name"`
	Properties HeadingProperties `json:"properties"`
}

// SetID implements the Module interface.
func (h *Heading) SetID(id string) {
	h.ID = id
}

// GetType implements the Module interface.
func (h *Heading) GetType() string {
	return h.Type
}

// Validate implements the Module interface.
func (h *Heading) Validate() error {
	// Check type.
	if err := ValidateType(h.Type); err != nil {
		return err
	}

	return nil
}

// HeadingProperties defines the header module properties.
type HeadingProperties struct {
	Title             string `json:"title"`
	Sublabel          string `json:"sublabel"`
	Size              string `json:"size"`
	Alignment         string `json:"alignment"`
	ImageAlignment    string `json:"image_alignment"`
	VerticalAlignment string `json:"vertical_alignment"`
	// Image string `json:"image"`
	ImageWidth int `json:"image_width"`
}
