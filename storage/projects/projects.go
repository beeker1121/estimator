package projects

// Database defines the projects database interface.
type Database interface {
	Create(p *Project) (*Project, error)
	GetByID(id string) (*Project, error)
	UpdateByID(id string, up *UpdateParams) (*Project, error)
}

// Project defines a project.
type Project struct {
	ID        string
	AccountID string
	Name      string
}

// UpdateParams defines the update parameters.
type UpdateParams struct {
	ID        *string
	AccountID *string
	Name      *string
}
