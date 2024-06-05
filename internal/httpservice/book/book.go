package book

import (
	"online-book-store/internal/model"

	"github.com/gofiber/fiber/v2"
)

func (h *BookHandler) GetBooks(c *fiber.Ctx) error {
	books, err := h.BookService.GetBooks(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(model.ErrorResponse{
			StatusCode: fiber.StatusInternalServerError,
			Message:    err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(model.JsonResponse{
		StatusCode: fiber.StatusOK,
		Message:    "success get books",
		Data:       books,
	})
}

func (h *BookHandler) GetBook(c *fiber.Ctx) error {
	book, err := h.BookService.GetBook(c.Context(), c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(model.ErrorResponse{
			StatusCode: fiber.StatusInternalServerError,
			Message:    err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(model.JsonResponse{
		StatusCode: fiber.StatusOK,
		Message:    "success get book",
		Data:       book,
	})
}
