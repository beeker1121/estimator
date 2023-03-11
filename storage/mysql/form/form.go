package form

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"

	"estimator/storage/form"

	"github.com/google/uuid"
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
	// insert a new form into the database.
	stmtInsert = `
INSERT INTO forms (id, modules)
VALUES (?, ?)
`
)

// Create creates a new form.
func (db *Database) Create(f *form.Form) (*form.Form, error) {
	// Map to local Form type.
	lf := &Form{
		ID: uuid.NewString(),
		Modules: Modules{
			Data: f.Modules,
		},
	}

	// Execute the query.
	if _, err := db.db.Exec(stmtInsert, lf.ID, lf.Modules); err != nil {
		return nil, err
	}

	// Set ID.
	f.ID = lf.ID

	return f, nil
}

// GetByID gets a form by the given ID.
func (db *Database) GetByID(id string) (*form.Form, error) {
	return &form.Form{}, nil
}

// Form defines a form.
type Form struct {
	ID      string
	Modules Modules
}

// Modules defines form modules.
type Modules struct {
	Data interface{}
}

// Value implements the driver interface.
func (m Modules) Value() (driver.Value, error) {
	b, err := json.Marshal(m.Data)
	if err != nil {
		return nil, err
	}

	return driver.Value(b), nil
}
