package order

import (
	"context"
	"online-book-store/internal/model"

	uuid "github.com/satori/go.uuid"
)

func (s *orderService) GetOrders(ctx context.Context, userID string) ([]model.Order, error) {
	orders, err := s.OrderRepo.GetOrders(ctx, userID)
	if err != nil {
		return []model.Order{}, err
	}

	for i := range orders {
		orders[i].CalculateTotalPrice()
	}

	return orders, err
}

func (s *orderService) GetOrder(ctx context.Context, userID, orderNo string) (*model.Order, error) {
	order, err := s.OrderRepo.GetOrder(ctx, userID, orderNo)
	if err != nil {
		return nil, err
	}

	order.CalculateTotalPrice()

	return order, nil
}

func (s *orderService) CreateOrder(ctx context.Context, userID string, order *model.Order) (*model.Order, error) {
	order.UserID = uuid.FromStringOrNil(userID)
	order.Status = "pending"

	for i, item := range order.Items {
		book, err := s.BookRepo.GetBook(ctx, item.BookID.String())
		if err != nil {
			return nil, err
		}
		order.Items[i].Price = book.Price
	}

	order, err := s.OrderRepo.CreateOrder(ctx, order)
	if err != nil {
		return nil, err
	}

	order.CalculateTotalPrice()

	return order, nil
}
