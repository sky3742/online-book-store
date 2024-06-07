package order

import (
	"context"
	"online-book-store/internal/constant"
	"online-book-store/internal/model"

	"github.com/gofiber/fiber/v2/log"
	uuid "github.com/satori/go.uuid"
)

func (s *orderService) GetOrders(ctx context.Context, userID string) ([]model.Order, *model.AppError) {
	orders, err := s.OrderRepo.GetOrders(ctx, userID)
	if err != nil {
		log.Errorf("error get orders: %v", err)
		return []model.Order{}, constant.ErrGetOrders
	}

	for i := range orders {
		orders[i].CalculateTotalPrice()
	}

	return orders, nil
}

func (s *orderService) GetOrder(ctx context.Context, userID, orderNo string) (*model.Order, *model.AppError) {
	order, err := s.OrderRepo.GetOrder(ctx, userID, orderNo)
	if err != nil {
		log.Errorf("error get order: %v", err)
		return nil, constant.ErrOrderNotFound
	}

	order.CalculateTotalPrice()

	return order, nil
}

func (s *orderService) CreateOrder(ctx context.Context, userID string, order *model.Order) (*model.Order, *model.AppError) {
	order.UserID = uuid.FromStringOrNil(userID)
	order.Status = "pending"

	for i, item := range order.Items {
		book, err := s.BookRepo.GetBook(ctx, item.BookID.String())
		if err != nil {
			log.Errorf("error get book: %v", err)
			return nil, constant.ErrGetBooks
		}
		order.Items[i].Price = book.Price
	}

	order, err := s.OrderRepo.CreateOrder(ctx, order)
	if err != nil {
		log.Errorf("error create order: %v", err)
		return nil, constant.ErrCreateOrder
	}

	order.CalculateTotalPrice()

	return order, nil
}
