package order

import (
	"online-book-store/internal/service"

	"github.com/gofiber/fiber/v2"
)

type OrderHandler struct {
	OrderService service.Order
}

type ConfigHandler struct {
	OrderService service.Order
}

func NewHandler(cfg ConfigHandler) *OrderHandler {
	return &OrderHandler{
		OrderService: cfg.OrderService,
	}
}

func (h *OrderHandler) SetPublicRoute(app fiber.Router) {

}

func (h *OrderHandler) SetPrivateRoute(app fiber.Router, authMiddleware fiber.Handler) {
	app.Use(authMiddleware)

	app.Get("", h.GetOrders)
	app.Get("/:orderNo", h.GetOrder)
	app.Post("", h.CreateOrder)
}
