package types

// Form defines a form.
type Form struct {
	ID         string         `json:"id"`
	ProjectID  string         `json:"project_id"`
	Name       string         `json:"name"`
	Properties FormProperties `json:"properties"`
	Button     FormButton     `json:"button"`
	Modules    []Module       `json:"modules"`
}

// FormProperties defines the form properties.
type FormProperties struct {
	BackgroundColor string `json:"background_color"`
	FontColor       string `json:"font_color"`
}

// FormButton defines the form button.
type FormButton struct {
	BackgroundColor string `json:"background_color"`
	Color           string `json:"color"`
	FontSize        string `json:"font_size"`
	FontFamily      string `json:"font_family"`
}
