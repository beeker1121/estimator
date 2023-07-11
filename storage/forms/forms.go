package forms

// Database defines the form database interface.
type Database interface {
	Create(f *Form) (*Form, error)
	GetByID(id string) (*Form, error)
	UpdateByID(id string, up *UpdateParams) (*Form, error)
}

// Form defines a form.
type Form struct {
	ID         string
	ProjectID  string
	Name       string
	Properties interface{}
	Button     interface{}
	Modules    interface{}
}

// UpdateParams defines the update parameters.
type UpdateParams struct {
	ID         *string
	ProjectID  *string
	Name       *string
	Properties *interface{}
	Button     *interface{}
	Modules    *interface{}
}
