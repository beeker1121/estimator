package types

// User defines a user.
type User struct {
	ID        string `json:"id"`
	AccountID string `json:"account_id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Role      string `json:"role"`
}

// UserUpdateParams defines the update parameters for users.
type UserUpdateParams struct {
	ID        *string `json:"id"`
	AccountID *string `json:"account_id"`
	Name      *string `json:"name"`
	Email     *string `json:"email"`
	Password  *string `json:"password"`
	Role      *string `json:"role"`
}
