package book

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// Service - the struct for our comment service
type Service struct {
	DB *gorm.DB
}

// NewService - returns a new book service
func NewService(db *gorm.DB) *Service {
	return &Service{
		DB: db,
	}
}

type Book struct {
	gorm.Model
	ISIN   int    `json:"ISIN"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Rating int    `json:"rating"`
}

func (s *Service) GetBooks(c *fiber.Ctx) error {
	var books []Book
	s.DB.Find(&books)
	return c.JSON(books)
}

func (s *Service) GetBook(c *fiber.Ctx) error {
	id := c.Params("id")
	var book Book
	s.DB.Find(&book, id)
	return c.JSON(book)
}

func (s *Service) NewBook(c *fiber.Ctx) error {
	book := new(Book)
	if err := c.BodyParser(book); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	s.DB.Create(&book)
	return c.JSON(book)
}

func (s *Service) DeleteBook(c *fiber.Ctx) error {
	ISIN := c.Params("isin")

	fmt.Println(ISIN)

	var book Book
	s.DB.First(&book, "ISIN = ?", ISIN)
	if book.Title == "" {
		return c.Status(500).SendString("No book found with given ID")
	}
	s.DB.Delete(&book)
	return c.SendString("Book successfully deleted")
}
