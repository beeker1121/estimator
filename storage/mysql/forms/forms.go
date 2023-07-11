package forms

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"

	"estimator/storage/forms"
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
INSERT INTO forms (id, project_id, name, properties, button, modules)
VALUES (?, ?, ?, ?, ?, ?)
`

	// stmtGetByID defines the SQL statement to
	// get a form from the database.
	stmtGetByID = `
SELECT * FROM forms
WHERE id=?
`

	// stmtUpdateByID defines the SQL statement
	// to update a form by the given ID.
	stmtUpdateByID = `
UPDATE forms
SET %s
WHERE id=?
`
)

// Form defines a form.
type Form struct {
	ID         string
	ProjectID  string
	Name       string
	Properties Properties
	Button     Button
	Modules    Modules
}

// Properties defines form properties.
type Properties struct {
	Data interface{}
}

// Value implements the driver interface.
func (p Properties) Value() (driver.Value, error) {
	b, err := json.Marshal(p.Data)
	if err != nil {
		return nil, err
	}

	return driver.Value(b), nil
}

// Scan implements the Scanner interface.
func (p *Properties) Scan(src any) error {
	val := src.([]uint8)
	return json.Unmarshal(val, &p.Data)
}

// Button defines form button.
type Button struct {
	Data interface{}
}

// Value implements the driver interface.
func (b Button) Value() (driver.Value, error) {
	j, err := json.Marshal(b.Data)
	if err != nil {
		return nil, err
	}

	return driver.Value(j), nil
}

// Scan implements the Scanner interface.
func (b *Button) Scan(src any) error {
	val := src.([]uint8)
	return json.Unmarshal(val, &b.Data)
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
func (db *Database) Create(f *forms.Form) (*forms.Form, error) {
	// Map to local Form type.
	lf := &Form{
		ID:        f.ID,
		ProjectID: f.ProjectID,
		Name:      f.Name,
		Properties: Properties{
			Data: f.Properties,
		},
		Button: Button{
			Data: f.Button,
		},
		Modules: Modules{
			Data: f.Modules,
		},
	}

	// Execute the query.
	if _, err := db.db.Exec(stmtInsert, lf.ID, lf.ProjectID, lf.Name, lf.Properties, lf.Button, lf.Modules); err != nil {
		return nil, err
	}

	return f, nil
}

// GetByID gets a form by the given ID.
func (db *Database) GetByID(id string) (*forms.Form, error) {
	// Create a new Form.
	f := &Form{
		Properties: Properties{},
		Button:     Button{},
		Modules:    Modules{},
	}

	// Execute the query.
	row := db.db.QueryRow(stmtGetByID, id)

	// Map columns to form.
	err := row.Scan(&f.ID, &f.ProjectID, &f.Name, &f.Properties, &f.Button, &f.Modules)
	switch {
	case err == sql.ErrNoRows:
		return nil, forms.ErrFormNotFound
	case err != nil:
		return nil, err
	}

	// Map to storage form type.
	gf := &forms.Form{
		ID:         f.ID,
		ProjectID:  f.ProjectID,
		Name:       f.Name,
		Properties: f.Properties.Data,
		Button:     f.Button.Data,
		Modules:    f.Modules.Data,
	}

	return gf, nil
}

// UpdateByID a form by the given ID.
func (db *Database) UpdateByID(id string, up *forms.UpdateParams) (*forms.Form, error) {
	// Create variables to hold the query fields
	// being updated and their new values.
	var queryFields string
	var queryValues []interface{}

	// Handle project ID field.
	if up.ProjectID != nil {
		if queryFields == "" {
			queryFields = "project_id=?"
		} else {
			queryFields += ", project_id=?"
		}

		queryValues = append(queryValues, *up.ProjectID)
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

	// Handle properties field.
	if up.Properties != nil {
		if queryFields == "" {
			queryFields = "properties=?"
		} else {
			queryFields += ", properties=?"
		}

		// Create new properties.
		p := Properties{
			Data: *up.Properties,
		}

		queryValues = append(queryValues, p)
	}

	// Handle button field.
	if up.Button != nil {
		if queryFields == "" {
			queryFields = "button=?"
		} else {
			queryFields += ", button=?"
		}

		// Create a new button.
		b := Button{
			Data: *up.Button,
		}

		queryValues = append(queryValues, b)
	}

	// Handle modules field.
	if up.Modules != nil {
		if queryFields == "" {
			queryFields = "modules=?"
		} else {
			queryFields += ", modules=?"
		}

		// Create new modules.
		m := Modules{
			Data: *up.Modules,
		}

		queryValues = append(queryValues, m)
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
	// form. Anything more complicated should use the
	// original statement constants.
	return db.GetByID(id)
}
