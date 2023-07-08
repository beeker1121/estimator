package accounts

import (
	"database/sql"

	"estimator/storage/accounts"
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
	// insert a new account into the database.
	stmtInsert = `
INSERT INTO accounts (id, name)
VALUES (?, ?)
`

	// stmtGetByID defines the SQL statement to
	// get an account from the database.
	stmtGetByID = `
SELECT * FROM accounts
WHERE id=?
`

	// stmtUpdateByID defines the SQL statement
	// to update an account by the given ID.
	stmtUpdateByID = `
UPDATE accounts
SET name=?
WHERE id=?
`
)

// Account defines an account.
type Account struct {
	ID   string
	Name string
}

// Create creates a new account.
func (db *Database) Create(a *accounts.Account) (*accounts.Account, error) {
	// Map to local Account type.
	la := &Account{
		ID:   a.ID,
		Name: a.Name,
	}

	// Execute the query.
	if _, err := db.db.Exec(stmtInsert, la.ID, la.Name); err != nil {
		return nil, err
	}

	return a, nil
}

// GetByID gets an account by the given ID.
func (db *Database) GetByID(id string) (*accounts.Account, error) {
	// Create a new Account.
	a := &Account{}

	// Execute the query.
	row := db.db.QueryRow(stmtGetByID, id)

	// Map columns to account.
	err := row.Scan(&a.ID, &a.Name)
	switch {
	case err == sql.ErrNoRows:
		return nil, accounts.ErrAccountNotFound
	case err != nil:
		return nil, err
	}

	// Map to storage account type.
	ga := &accounts.Account{
		ID:   a.ID,
		Name: a.Name,
	}

	return ga, nil
}

// UpdateByID a form by the given ID.
func (db *Database) UpdateByID(id string, a *accounts.Account) (*accounts.Account, error) {
	// Map to local Account type.
	la := &Account{
		ID:   id,
		Name: a.Name,
	}

	// Execute the query.
	if _, err := db.db.Exec(stmtUpdateByID, la.Name, la.ID); err != nil {
		return nil, err
	}

	return a, nil
}
