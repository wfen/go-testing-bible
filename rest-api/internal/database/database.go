package database

import (
	"errors"
	"fmt"

	"gorm.io/driver/sqlite"
	_ "gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	DBConn *gorm.DB
)

// NewDatabase -
func NewDatabase() (*gorm.DB, error) {
	var err error
	//DBConn, err = gorm.Open(sqlite.Open("rest-api/books.db"), &gorm.Config{})
	DBConn, err = gorm.Open(sqlite.Open("books.db"), &gorm.Config{})
	if err != nil {
		return nil, errors.New("Failed to connect to database")
	}
	fmt.Println("Database connection successfully opened")
	return DBConn, nil
}
