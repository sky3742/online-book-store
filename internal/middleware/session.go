package middleware

import (
	"online-book-store/internal/constant"
	"online-book-store/internal/utils"

	"github.com/gofiber/fiber/v2"
)

func SessionMiddleware(c *fiber.Ctx) error {
	token := c.Get("X-Session-Token")
	if token == "" {
		return utils.ErrorResponse(c, constant.ErrMissingSession)
	}

	userSession, errx := utils.GetSession(token)
	if errx != nil {
		return utils.ErrorResponse(c, errx)
	}

	c.Locals("token", token)
	c.Locals("userID", userSession.ID)

	return c.Next()
}
