package transport

import (
	"github.com/gofiber/fiber/v2"

	"github.com/wfen/go-testing-bible/rest-api/internal/book"
)

// Fiber uses fasthttp. Everything that is built on top of fasthttp should not be
// considered unless you actually know what you are doing. Mainly because fasthttp
// does not support http/2. Fiber also feels different from the standard library.

// Handler - stores pointer to our comments service
type Handler struct {
	App *fiber.App
	Svc *book.Service
}

// NewHandler - returns a pointer to a Handler
func NewHandler(service *book.Service) *Handler {
	return &Handler{
		App: fiber.New(),
		Svc: service,
	}
}

func helloWorld(c *fiber.Ctx) error {
	c.Send([]byte("Hello, World!"))
	return nil
}

func (h Handler) setupRoutes() {
	h.App.Get("/", helloWorld)
	h.App.Get("/api/v1/book", h.Svc.GetBooks)
	h.App.Get("/api/v1/book/:id", h.Svc.GetBook)
	h.App.Post("/api/v1/book", h.Svc.NewBook)
	h.App.Delete("/api/v1/book/:isin", h.Svc.DeleteBook)
}

// Setup - set's up our fiber app and the routes
// returns a pointer to app
func Setup(svc *book.Service) *Handler {
	handler := NewHandler(svc)
	handler.setupRoutes()
	return handler
}
