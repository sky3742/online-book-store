package account

import (
	"online-book-store/internal/service"

	"github.com/gofiber/fiber/v2"
)

type AccountHandler struct {
	AccountService service.Account
}

type ConfigHandler struct {
	AccountService service.Account
}

func NewHandler(cfg ConfigHandler) *AccountHandler {
	return &AccountHandler{
		AccountService: cfg.AccountService,
	}
}

func (h *AccountHandler) SetPublicRoute(app fiber.Router) {
	app.Post("/register", h.RegisterAccount)
	app.Post("/login", h.LoginAccount)
}

func (h *AccountHandler) SetPrivateRoute(app fiber.Router, authMiddleware fiber.Handler) {
	logoutGroup := app.Group("/logout")
	logoutGroup.Use(authMiddleware)
	logoutGroup.Post("", h.LogoutAccount)
}
