package main

import (
	"online-book-store/cmd/internal"
	"online-book-store/internal/middleware"
	"online-book-store/internal/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	utils.InitEnv()

	server := fiber.New()
	server.Use(logger.New())
	server.Use(healthcheck.New())
	server.Use(recover.New())

	app := server.Group("/api/v1")

	service := internal.InitServices()
	httpService := internal.InitHttpService(service)

	httpService.Account.SetPublicRoute(app)
	httpService.Account.SetPrivateRoute(app, middleware.SessionMiddleware)

	httpService.Book.SetPublicRoute(app.Group("/books"))

	httpService.Order.SetPrivateRoute(app.Group("/orders"), middleware.SessionMiddleware)

	server.Listen(":3000")
}
