package middleware

import (
	"online-book-store/internal/model"
	"online-book-store/internal/utils"

	"github.com/gofiber/fiber/v2"
)

func SessionMiddleware(c *fiber.Ctx) error {
	token := c.Get("X-Session-Token")

	userSession, err := utils.GetSession(token)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(model.ErrorResponse{
			StatusCode: fiber.StatusUnauthorized,
			Message:    err.Error(),
		})
	}

	c.Locals("token", token)
	c.Locals("userID", userSession.ID)

	return c.Next()
}
