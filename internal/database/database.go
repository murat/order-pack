package database

import (
	"errors"
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Database struct {
	Conn *gorm.DB
}

func New(dbName string) (*Database, error) {
	if dbName == "" {
		return nil, errors.New("dbName is not provided")
	}

	db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("could not open database, %w", err)
	}

	return &Database{Conn: db}, nil
}
