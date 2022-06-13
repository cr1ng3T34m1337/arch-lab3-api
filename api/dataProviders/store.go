package dataProviders

import (
	"database/sql"
)

type Store struct {
	Db *sql.DB
}

func StoreProvider(db *sql.DB) *Store {
	return &Store{Db: db}
}
