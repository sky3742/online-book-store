package book

import (
	"online-book-store/internal/model"
	"online-book-store/internal/utils"

	"github.com/gofiber/fiber/v2"
)

func (h *BookHandler) GetBooks(c *fiber.Ctx) error {
	books, errx := h.BookService.GetBooks(c.Context())
	if errx != nil {
		return utils.ErrorResponse(c, errx)
	}

	return c.Status(fiber.StatusOK).JSON(model.JsonResponse{
		StatusCode: fiber.StatusOK,
		Message:    "success get books",
		Data:       books,
	})
}

func (h *BookHandler) GetBook(c *fiber.Ctx) error {
	book, errx := h.BookService.GetBook(c.Context(), c.Params("id"))
	if errx != nil {
		return utils.ErrorResponse(c, errx)
	}

	return c.Status(fiber.StatusOK).JSON(model.JsonResponse{
		StatusCode: fiber.StatusOK,
		Message:    "success get book",
		Data:       book,
	})
}
