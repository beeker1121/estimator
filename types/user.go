package types

// User defines a user.
type User struct {
	ID        string `json:"id"`
	AccountID string `json:"account_id"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

// UserUpdateParams defines the parameters for the Update method.
type UserUpdateParams struct {
	ID        *string `json:"id"`
	AccountID *string `json:"account_id"`
	Email     *string `json:"email"`
	Password  *string `json:"password"`
}
