package order

import (
	"online-book-store/internal/model"

	"github.com/gofiber/fiber/v2"
)

func (h *OrderHandler) GetOrders(c *fiber.Ctx) error {
	userID := c.Locals("userID").(string)

	orders, err := h.OrderService.GetOrders(c.Context(), userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(model.ErrorResponse{
			StatusCode: fiber.StatusInternalServerError,
			Message:    err.Error(),
		})
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

	data, err := h.OrderService.GetOrder(c.Context(), userID, orderNo)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(model.ErrorResponse{
			StatusCode: fiber.StatusInternalServerError,
			Message:    err.Error(),
		})
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
		return c.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse{
			StatusCode: fiber.StatusBadRequest,
			Message:    err.Error(),
		})
	}

	data, err := h.OrderService.CreateOrder(c.Context(), userID, &order)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(model.ErrorResponse{
			StatusCode: fiber.StatusInternalServerError,
			Message:    err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(model.JsonResponse{
		StatusCode: fiber.StatusOK,
		Message:    "success create order",
		Data:       data,
	})
}
