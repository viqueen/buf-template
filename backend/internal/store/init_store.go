package store

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitStore() (*gorm.DB, error) {
	connectionString := "postgres://todo:authz@localhost:5432/todo?sslmode=disable" //nolint: gosec

	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	return db, nil
}
