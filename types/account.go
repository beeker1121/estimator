package types

// Account defines an account.
type Account struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// AccountUpdateParams defines the update parameters for accounts.
type AccountUpdateParams struct {
	ID   *string `json:"id"`
	Name *string `json:"name"`
}
