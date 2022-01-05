package database

import (
	"gorm.io/gorm"

	"github.com/wfen/go-testing-bible/rest-api/internal/book"
)

// MigrateDB - migrates our database and creates our book table
func MigrateDB(db *gorm.DB) error {
	if err := db.AutoMigrate(&book.Book{}); err != nil {
		return err
	}
	return nil
}
