package projects

// Database defines the projects database interface.
type Database interface {
	Create(p *Project) (*Project, error)
	GetByID(id string) (*Project, error)
	UpdateByID(id string, up *UpdateParams) (*Project, error)
	Get(gp *GetParams) (*Projects, error)
}

// Project defines a project.
type Project struct {
	ID        string
	AccountID string
	Name      string
}

// Projects defines a set of projects.
type Projects struct {
	Projects []*Project `json:"projects"`
	Total    int        `json:"total"`
}

// UpdateParams defines the update parameters.
type UpdateParams struct {
	ID        *string
	AccountID *string
	Name      *string
}

// GetParams defines the get parameters.
type GetParams struct {
	ID        *string
	AccountID *string
	Name      *string
	Offset    int
	Limit     int
}
