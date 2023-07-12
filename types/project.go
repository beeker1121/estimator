package types

// Project defines an project.
type Project struct {
	ID        string `json:"id"`
	AccountID string `json:"account_id"`
	Name      string `json:"name"`
}

// ProjectUpdateParams defines the update parameters for projects.
type ProjectUpdateParams struct {
	ID        *string `json:"id"`
	AccountID *string `json:"account_id"`
	Name      *string `json:"name"`
}
