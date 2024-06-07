package order

import (
	"online-book-store/internal/constant"
	"online-book-store/internal/model"
	"online-book-store/internal/utils"

	"github.com/gofiber/fiber/v2"
)

func (h *OrderHandler) GetOrders(c *fiber.Ctx) error {
	userID := c.Locals("userID").(string)

	orders, errx := h.OrderService.GetOrders(c.Context(), userID)
	if errx != nil {
		return utils.ErrorResponse(c, errx)
	}

	return c.Status(fiber.StatusOK).JSON(model.JsonResponse{
		StatusCode: fiber.StatusOK,
		Message:    "success get orders",
		Data:       orders,
	})
}

func (h *OrderHandler) GetOrder(c *fiber.Ctx) error {
	userID := c.Locals("userID").(string)

	orderNo := c.Params("orderNo")

	data, errx := h.OrderService.GetOrder(c.Context(), userID, orderNo)
	if errx != nil {
		return utils.ErrorResponse(c, errx)
	}

	return c.Status(fiber.StatusOK).JSON(model.JsonResponse{
		StatusCode: fiber.StatusOK,
		Message:    "success get order",
		Data:       data,
	})
}

func (h *OrderHandler) CreateOrder(c *fiber.Ctx) error {
	userID := c.Locals("userID").(string)

	var order model.Order

	err := c.BodyParser(&order)
	if err != nil {
		c.Context().Logger().Printf("error parsing body: %v", err)
		return utils.ErrorResponse(c, constant.ErrBadRequest)
	}

	data, errx := h.OrderService.CreateOrder(c.Context(), userID, &order)
	if errx != nil {
		return utils.ErrorResponse(c, errx)
	}

	return c.Status(fiber.StatusCreated).JSON(model.JsonResponse{
		StatusCode: fiber.StatusCreated,
		Message:    "success create order",
		Data:       data,
	})
}
