package database

import (
	"errors"
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func New(path string) (*gorm.DB, error) {
	if path == "" {
		return nil, errors.New("path is not provided")
	}

	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("could not open database, %w", err)
	}

	return db, nil
}
