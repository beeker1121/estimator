package form

import (
	"database/sql"

	"estimator/storage/form"
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

// Create creates a new form.
func (db *Database) Create(f *form.Form) (*form.Form, error) {
	return &form.Form{}, nil
}

// GetByID gets a form by the given ID.
func (db *Database) GetByID(id string) (*form.Form, error) {
	return &form.Form{}, nil
}
