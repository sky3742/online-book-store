package account

import (
	"online-book-store/internal/model"
	"online-book-store/internal/utils"

	"github.com/gofiber/fiber/v2"
)

func (h *AccountHandler) CreateAccount(c *fiber.Ctx) error {
	var account model.Account

	err := c.BodyParser(&account)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse{
			StatusCode: fiber.StatusBadRequest,
			Message:    err.Error(),
		})
	}

	acc, err := h.AccountService.CreateAccount(c.Context(), &account)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(model.ErrorResponse{
			StatusCode: fiber.StatusInternalServerError,
			Message:    err.Error(),
		})
	}

	token := utils.CreateSession(acc.ID.String())

	return c.Status(fiber.StatusOK).JSON(model.JsonResponse{
		StatusCode: fiber.StatusOK,
		Message:    "success create account",
		Data: map[string]string{
			"token": token,
		},
	})
}

func (h *AccountHandler) LoginAccount(c *fiber.Ctx) error {
	var account model.Account

	err := c.BodyParser(&account)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse{
			StatusCode: fiber.StatusBadRequest,
			Message:    err.Error(),
		})
	}

	acc, err := h.AccountService.GetAccount(c.Context(), &account)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(model.ErrorResponse{
			StatusCode: fiber.StatusInternalServerError,
			Message:    err.Error(),
		})
	}

	token := utils.CreateSession(acc.ID.String())

	return c.Status(fiber.StatusOK).JSON(model.JsonResponse{
		StatusCode: fiber.StatusOK,
		Message:    "success login account",
		Data: map[string]string{
			"token": token,
		},
	})
}

func (h *AccountHandler) LogoutAccount(c *fiber.Ctx) error {
	token := c.Locals("token").(string)

	utils.RemoveSession(token)

	return c.Status(fiber.StatusOK).JSON(model.JsonResponse{
		StatusCode: fiber.StatusOK,
		Message:    "success logout account",
	})
}
