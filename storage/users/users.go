package users

// Database defines the users database interface.
type Database interface {
	Create(u *User) (*User, error)
	GetByID(id string) (*User, error)
	GetByIDAndAccountID(id, accountID string) (*User, error)
	GetByEmail(email string) (*User, error)
	UpdateByID(id string, up *UpdateParams) (*User, error)
}

// User defines a user.
type User struct {
	ID        string
	AccountID string
	Email     string
	Password  string
}

// UpdateParams defines the parameters for the Update method.
type UpdateParams struct {
	ID        *string
	AccountID *string
	Email     *string
	Password  *string
}
