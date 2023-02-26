package form

// Database defines the form database interface.
type Database interface {
	Create(f *Form) (*Form, error)
	GetByID(id string) (*Form, error)
}

// Form defines a form.
type Form struct {
	ID string
}
