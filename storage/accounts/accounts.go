package accounts

// Database defines the accounts database interface.
type Database interface {
	Create(a *Account) (*Account, error)
	GetByID(id string) (*Account, error)
	UpdateByID(id string, a *Account) (*Account, error)
}

// Account defines an account.
type Account struct {
	ID   string
	Name string
}
