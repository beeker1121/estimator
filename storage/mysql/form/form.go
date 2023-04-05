package form

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"

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

const (
	// stmtInsert defines the SQL statement to
	// insert a new form into the database.
	stmtInsert = `
INSERT INTO forms (id, modules)
VALUES (?, ?)
`

	// stmtGetByID defines the SQL statement to
	// get a form from the datbase.
	stmtGetByID = `
SELECT * FROM forms
WHERE id=?
`

	// stmtUpdateByID defines the SQL statement
	// to update a form by the given ID.
	stmtUpdateByID = `
UPDATE forms
SET modules=?
WHERE id=?
`
)

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

// Scan implements the Scanner interface.
func (m *Modules) Scan(src any) error {
	val := src.([]uint8)
	return json.Unmarshal(val, &m.Data)
}

// Create creates a new form.
func (db *Database) Create(f *form.Form) (*form.Form, error) {
	// Map to local Form type.
	lf := &Form{
		ID: f.ID,
		Modules: Modules{
			Data: f.Modules,
		},
	}

	// Execute the query.
	if _, err := db.db.Exec(stmtInsert, lf.ID, lf.Modules); err != nil {
		return nil, err
	}

	return f, nil
}

// GetByID gets a form by the given ID.
func (db *Database) GetByID(id string) (*form.Form, error) {
	// Create a new Form.
	f := &Form{
		Modules: Modules{},
	}

	// Execute the query.
	row := db.db.QueryRow(stmtGetByID, id)

	// Map columns to form.
	err := row.Scan(&f.ID, &f.Modules)
	switch {
	case err == sql.ErrNoRows:
		return nil, form.ErrFormNotFound
	case err != nil:
		return nil, err
	}

	// Map to storage form type.
	gf := &form.Form{
		ID:      f.ID,
		Modules: f.Modules.Data,
	}

	return gf, nil
}

// UpdateByID a form by the given ID.
func (db *Database) UpdateByID(id string, f *form.Form) (*form.Form, error) {
	// Map to local Form type.
	lf := &Form{
		ID: id,
		Modules: Modules{
			Data: f.Modules,
		},
	}

	// Execute the query.
	if _, err := db.db.Exec(stmtUpdateByID, lf.Modules, lf.ID); err != nil {
		return nil, err
	}

	return f, nil
}
