package users

import (
	"database/sql"
	"fmt"

	"estimator/storage/users"
)

// Database defines the database.
type Database struct {
	db *sql.DB
}

// New creates a new database.
func New(db *sql.DB) *Database {
	return &Database{
		db: db,
	}
}

const (
	// stmtInsert defines the SQL statement to
	// insert a new user into the database.
	stmtInsert = `
INSERT INTO users (id, account_id, name, email, password, role)
VALUES (?, ?, ?, ?, ?, ?)
`

	// stmtGetByID defines the SQL statement to
	// get a user from the database by ID.
	stmtGetByID = `
SELECT * FROM users
WHERE id=?
`

	// stmtGetByIDAndAccountID defines the SQL statement
	// to get a user from the database by ID and account ID.
	stmtGetByIDAndAccountID = `
SELECT * FROM users
WHERE id=?
AND account_id=?
`

	// stmtGetByEmail defines the SQL statement to
	// get a user from the database by email.
	stmtGetByEmail = `
SELECT * FROM users
WHERE email=?
`

	// stmtUpdateByID defines the SQL statement
	// to update a user by the given ID.
	stmtUpdateByID = `
UPDATE users
SET %s
WHERE id=?
`
)

// User defines a user.
type User struct {
	ID        string
	AccountID string
	Name      string
	Email     string
	Password  string
	Role      string
}

// Create creates a new user.
func (db *Database) Create(u *users.User) (*users.User, error) {
	// Map to local User type.
	lu := &User{
		ID:        u.ID,
		AccountID: u.AccountID,
		Name:      u.Name,
		Email:     u.Email,
		Password:  u.Password,
		Role:      u.Role,
	}

	// Execute the query.
	if _, err := db.db.Exec(stmtInsert, lu.ID, lu.AccountID, lu.Name, lu.Email, lu.Password, lu.Role); err != nil {
		return nil, err
	}

	return u, nil
}

// GetByID gets a user by the given ID.
func (db *Database) GetByID(id string) (*users.User, error) {
	// Create a new User.
	u := &User{}

	// Execute the query.
	row := db.db.QueryRow(stmtGetByID, id)

	// Map columns to user.
	err := row.Scan(&u.ID, &u.AccountID, &u.Name, &u.Email, &u.Password, &u.Role)
	switch {
	case err == sql.ErrNoRows:
		return nil, users.ErrUserNotFound
	case err != nil:
		return nil, err
	}

	// Map to storage user type.
	su := &users.User{
		ID:        u.ID,
		AccountID: u.AccountID,
		Name:      u.Name,
		Email:     u.Email,
		Password:  u.Password,
		Role:      u.Role,
	}

	return su, nil
}

// GetByIDAndAccountID gets a user by the given ID and account ID.
func (db *Database) GetByIDAndAccountID(id, accountID string) (*users.User, error) {
	// Create a new User.
	u := &User{}

	// Execute the query.
	row := db.db.QueryRow(stmtGetByIDAndAccountID, id, accountID)

	// Map columns to user.
	err := row.Scan(&u.ID, &u.AccountID, &u.Name, &u.Email, &u.Password, &u.Role)
	switch {
	case err == sql.ErrNoRows:
		return nil, users.ErrUserNotFound
	case err != nil:
		return nil, err
	}

	// Map to storage user type.
	uf := &users.User{
		ID:        u.ID,
		AccountID: u.AccountID,
		Name:      u.Name,
		Email:     u.Email,
		Password:  u.Password,
		Role:      u.Role,
	}

	return uf, nil
}

// GetByEmail gets a user by the given email.
func (db *Database) GetByEmail(email string) (*users.User, error) {
	// Create a new User.
	u := &User{}

	// Execute the query.
	row := db.db.QueryRow(stmtGetByEmail, email)

	// Map columns to user.
	err := row.Scan(&u.ID, &u.AccountID, &u.Name, &u.Email, &u.Password, &u.Role)
	switch {
	case err == sql.ErrNoRows:
		return nil, users.ErrUserNotFound
	case err != nil:
		return nil, err
	}

	// Map to storage user type.
	uf := &users.User{
		ID:        u.ID,
		AccountID: u.AccountID,
		Name:      u.Name,
		Email:     u.Email,
		Password:  u.Password,
		Role:      u.Role,
	}

	return uf, nil
}

// UpdateByID a user by the given ID.
func (db *Database) UpdateByID(id string, up *users.UpdateParams) (*users.User, error) {
	// Create variables to hold the query fields
	// being updated and their new values.
	var queryFields string
	var queryValues []interface{}

	// Handle account ID field.
	if up.AccountID != nil {
		if queryFields == "" {
			queryFields = "account_id=?"
		} else {
			queryFields += ", account_id=?"
		}

		queryValues = append(queryValues, *up.AccountID)
	}

	// Handle name field.
	if up.Name != nil {
		if queryFields == "" {
			queryFields = "name=?"
		} else {
			queryFields += ", name=?"
		}

		queryValues = append(queryValues, *up.Name)
	}

	// Handle email field.
	if up.Email != nil {
		if queryFields == "" {
			queryFields = "email=?"
		} else {
			queryFields += ", email=?"
		}

		queryValues = append(queryValues, *up.Email)
	}

	// Handle password field.
	if up.Password != nil {
		if queryFields == "" {
			queryFields = "password=?"
		} else {
			queryFields += ", password=?"
		}

		queryValues = append(queryValues, *up.Password)
	}

	// Handle role field.
	if up.Role != nil {
		if queryFields == "" {
			queryFields = "role=?"
		} else {
			queryFields += ", role=?"
		}

		queryValues = append(queryValues, *up.Role)
	}

	// Check if query is empty.
	if queryFields == "" {
		return db.GetByID(id)
	}

	// Build the full query.
	query := fmt.Sprintf(stmtUpdateByID, queryFields)
	queryValues = append(queryValues, id)

	// Execute the query.
	_, err := db.db.Exec(query, queryValues...)
	if err != nil {
		return nil, err
	}

	// Since the GetByID method is straight forward,
	// we can use this method to retrieve the updated
	// todo. Anything more complicated should use the
	// original statement constants.
	return db.GetByID(id)
}
