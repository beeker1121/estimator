package types

// Project defines an project.
type Project struct {
	ID        string `json:"id"`
	AccountID string `json:"account_id"`
	Name      string `json:"name"`
}

// Projects defines a set of projects.
type Projects struct {
	Projects []*Project `json:"projects"`
	Total    int        `json:"total"`
}

// ProjectUpdateParams defines the update parameters for projects.
type ProjectUpdateParams struct {
	ID        *string `json:"id"`
	AccountID *string `json:"account_id"`
	Name      *string `json:"name"`
}

// ProjectGetParams defines the get parameters for projects.
type ProjectGetParams struct {
	ID        *string `json:"id"`
	AccountID *string `json:"account_id"`
	Name      *string `json:"name"`
	Offset    int
	Limit     int
}
