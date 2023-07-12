package projects

import (
	"database/sql"
	"fmt"

	"estimator/storage/projects"
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
	// insert a new project into the database.
	stmtInsert = `
INSERT INTO projects (id, account_id, name)
VALUES (?, ?, ?)
`

	// stmtGetByID defines the SQL statement to
	// get an project from the database.
	stmtGetByID = `
SELECT * FROM projects
WHERE id=?
`

	// stmtUpdateByID defines the SQL statement
	// to update an project by the given ID.
	stmtUpdateByID = `
UPDATE projects
SET %s
WHERE id=?
`
)

// Project defines an project.
type Project struct {
	ID        string
	AccountID string
	Name      string
}

// Create creates a new project.
func (db *Database) Create(p *projects.Project) (*projects.Project, error) {
	// Map to local Project type.
	lp := &Project{
		ID:        p.ID,
		AccountID: p.AccountID,
		Name:      p.Name,
	}

	// Execute the query.
	if _, err := db.db.Exec(stmtInsert, lp.ID, lp.AccountID, lp.Name); err != nil {
		return nil, err
	}

	return p, nil
}

// GetByID gets an project by the given ID.
func (db *Database) GetByID(id string) (*projects.Project, error) {
	// Create a new Project.
	p := &Project{}

	// Execute the query.
	row := db.db.QueryRow(stmtGetByID, id)

	// Map columns to project.
	err := row.Scan(&p.ID, &p.AccountID, &p.Name)
	switch {
	case err == sql.ErrNoRows:
		return nil, projects.ErrProjectNotFound
	case err != nil:
		return nil, err
	}

	// Map to storage project type.
	sp := &projects.Project{
		ID:   p.ID,
		Name: p.Name,
	}

	return sp, nil
}

// UpdateByID a form by the given ID.
func (db *Database) UpdateByID(id string, up *projects.UpdateParams) (*projects.Project, error) {
	// Create variables to hold the query fields
	// being updated and their new values.
	var queryFields string
	var queryValues []interface{}

	// Handle account ID.
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
	// project. Anything more complicated should use the
	// original statement constants.
	return db.GetByID(id)
}
