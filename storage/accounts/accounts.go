package accounts

// Database defines the accounts database interface.
type Database interface {
	Create(a *Account) (*Account, error)
	GetByID(id string) (*Account, error)
	UpdateByID(id string, up *UpdateParams) (*Account, error)
}

// Account defines an account.
type Account struct {
	ID   string
	Name string
}

// UpdateParams defines the update parameters.
type UpdateParams struct {
	ID   *string
	Name *string
}
