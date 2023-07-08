package mysql

import (
	"database/sql"

	"estimator/storage"
	"estimator/storage/mysql/accounts"
	"estimator/storage/mysql/forms"
	"estimator/storage/mysql/users"
)

// New returns a new storage implementation that uses MYSQL as the backend
// database.
func New(db *sql.DB) *storage.Storage {
	store := &storage.Storage{
		Accounts: accounts.New(db),
		Users:    users.New(db),
		Forms:    forms.New(db),
	}

	return store
}
