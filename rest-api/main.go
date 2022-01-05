package main

import (
	"log"

	_ "gorm.io/driver/sqlite"

	"github.com/wfen/go-testing-bible/rest-api/internal/book"
	"github.com/wfen/go-testing-bible/rest-api/internal/database"
	"github.com/wfen/go-testing-bible/rest-api/internal/transport"
)

func Run() error {
	var err error
	db, err := database.NewDatabase()
	if err != nil {
		return err
	}
	err = database.MigrateDB(db)
	if err != nil {
		return err
	}

	bookSvc := book.NewService(db)
	handler := transport.Setup(bookSvc)
	handler.App.Listen(":3000")
	return nil
}

func main() {
	if err := Run(); err != nil {
		log.Println("Error starting up our REST API")
		log.Fatal(err)
	}
}
