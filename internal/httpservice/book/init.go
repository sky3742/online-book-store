package book

import (
	"online-book-store/internal/service"

	"github.com/gofiber/fiber/v2"
)

type BookHandler struct {
	BookService service.Book
}

type ConfigHandler struct {
	BookService service.Book
}

func NewHandler(cfg ConfigHandler) *BookHandler {
	return &BookHandler{
		BookService: cfg.BookService,
	}
}

func (h *BookHandler) SetPublicRoute(app fiber.Router) {
	app.Get("/", h.GetBooks)
	app.Get("/:id", h.GetBook)
}
