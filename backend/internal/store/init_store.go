package store

import (
	"database/sql"
	"fmt"
)

func InitStore() (*sql.DB, error) {
	connectionString := "postgres://todo:authz@localhost:5432/todo?sslmode=disable" //nolint: gosec

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	return db, nil
}
