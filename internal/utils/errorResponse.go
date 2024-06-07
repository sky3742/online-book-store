package utils

import (
	"online-book-store/internal/model"

	"github.com/gofiber/fiber/v2"
)

func ErrorResponse(c *fiber.Ctx, err *model.AppError) error {
	return c.Status(err.HttpStatus).JSON(model.ErrorResponse{
		StatusCode: err.HttpStatus,
		ErrorCode:  err.Code,
		Message:    err.Message,
	})
}
