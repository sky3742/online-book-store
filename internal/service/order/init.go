package order

import (
	"online-book-store/internal/repository"
	"online-book-store/internal/service"
)

type orderService struct {
	BookRepo  repository.BookProvider
	OrderRepo repository.OrderProvider
}

type OrderServiceConfig struct {
	BookRepo  repository.BookProvider
	OrderRepo repository.OrderProvider
}

func NewOrderService(cfg OrderServiceConfig) service.Order {
	return &orderService{
		BookRepo:  cfg.BookRepo,
		OrderRepo: cfg.OrderRepo,
	}
}
