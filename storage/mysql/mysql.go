package mysql

import (
	"database/sql"

	"estimator/storage"
	"estimator/storage/mysql/form"
	"estimator/storage/mysql/users"
)

// New returns a new storage implementation that uses MYSQL as the backend
// database.
func New(db *sql.DB) *storage.Storage {
	store := &storage.Storage{
		Users: users.New(db),
		Form:  form.New(db),
	}

	return store
}
