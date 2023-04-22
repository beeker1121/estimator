package users

// Database defines the users database interface.
type Database interface {
	Create(u *User) (*User, error)
	GetByID(id string) (*User, error)
	GetByEmail(email string) (*User, error)
	UpdateByID(id string, u *User) (*User, error)
}

// User defines a user.
type User struct {
	ID       string
	Email    string
	Password string
}
