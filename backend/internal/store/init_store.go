package store

import "database/sql"

func InitStore() (*sql.DB, error) {
	connectionString := "postgres://todo:authz@localhost:5432/todo?sslmode=disable"
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}
	return db, nil
}
