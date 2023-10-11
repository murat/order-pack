package database

import (
	"errors"
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Conn struct {
	DB *gorm.DB
}

func New(dbName string) (*Conn, error) {
	if dbName == "" {
		return nil, errors.New("dbName is not provided")
	}

	db, err := gorm.Open(sqlite.Open(dbName+".db"), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("could not open database, %w", err)
	}

	return &Conn{DB: db}, nil
}
