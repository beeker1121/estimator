package users

import (
	"database/sql"

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
INSERT INTO users (id, email, password)
VALUES (?, ?, ?)
`

	// stmtGetByID defines the SQL statement to
	// get a user from the database.
	stmtGetByID = `
SELECT * FROM users
WHERE id=?
`

	// stmtUpdateByID defines the SQL statement
	// to update a user by the given ID.
	stmtUpdateByID = `
UPDATE users
SET email=?, password=?
WHERE id=?
`
)

// User defines a user.
type User struct {
	ID       string
	Email    string
	Password string
}

// Create creates a new user.
func (db *Database) Create(u *users.User) (*users.User, error) {
	// Map to local User type.
	lu := &User{
		ID:       u.ID,
		Email:    u.Email,
		Password: u.Password,
	}

	// Execute the query.
	if _, err := db.db.Exec(stmtInsert, lu.ID, lu.Email, lu.Password); err != nil {
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
	err := row.Scan(&u.ID, &u.Email, &u.Password)
	switch {
	case err == sql.ErrNoRows:
		return nil, users.ErrUserNotFound
	case err != nil:
		return nil, err
	}

	// Map to storage user type.
	uf := &users.User{
		ID:       u.ID,
		Email:    u.Email,
		Password: u.Password,
	}

	return uf, nil
}

// UpdateByID a user by the given ID.
func (db *Database) UpdateByID(id string, u *users.User) (*users.User, error) {
	// Map to local User type.
	lu := &User{
		ID:       id,
		Email:    u.Email,
		Password: u.Password,
	}

	// Execute the query.
	if _, err := db.db.Exec(stmtUpdateByID, lu.Email, lu.Password, lu.ID); err != nil {
		return nil, err
	}

	return u, nil
}
