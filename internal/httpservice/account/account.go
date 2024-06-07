package account

import (
	"online-book-store/internal/constant"
	"online-book-store/internal/model"
	"online-book-store/internal/utils"

	"github.com/gofiber/fiber/v2"
)

func (h *AccountHandler) RegisterAccount(c *fiber.Ctx) error {
	var account model.Account

	err := c.BodyParser(&account)
	if err != nil {
		c.Context().Logger().Printf("error parsing body: %v", err)
		return utils.ErrorResponse(c, constant.ErrBadRequest)
	}

	acc, errx := h.AccountService.RegisterAccount(c.Context(), &account)
	if errx != nil {
		return utils.ErrorResponse(c, errx)
	}

	token := utils.CreateSession(acc.ID.String())

	return c.Status(fiber.StatusCreated).JSON(model.JsonResponse{
		StatusCode: fiber.StatusCreated,
		Message:    "success register account",
		Data: map[string]string{
			"token": token,
		},
	})
}

func (h *AccountHandler) LoginAccount(c *fiber.Ctx) error {
	var account model.Account

	err := c.BodyParser(&account)
	if err != nil {
		c.Context().Logger().Printf("error parsing body: %v", err)
		return utils.ErrorResponse(c, constant.ErrBadRequest)
	}

	acc, errx := h.AccountService.GetAccount(c.Context(), &account)
	if errx != nil {
		return utils.ErrorResponse(c, errx)
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
