package form

// Database defines the form database interface.
type Database interface {
	Create(f *Form) (*Form, error)
	GetByID(id string) (*Form, error)
	UpdateByID(id string, f *Form) (*Form, error)
}

// Form defines a form.
type Form struct {
	ID      string
	Modules interface{}
}
