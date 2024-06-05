package order

import (
	"context"
	"online-book-store/internal/model"
)

func (r *orderRepo) GetOrders(ctx context.Context, userID string) ([]model.Order, error) {
	var orders []model.Order

	result := r.db.
		Model(&model.Order{}).
		Preload("Items").
		Where("user_id = ?", userID).
		Find(&orders)

	if result.Error != nil {
		return []model.Order{}, result.Error
	}

	return orders, nil
}

func (r *orderRepo) GetOrder(ctx context.Context, userID, orderID string) (*model.Order, error) {
	var order model.Order

	result := r.db.
		Model(&model.Order{}).
		Preload("Items").
		Where("id = ? AND user_id = ?", orderID, userID).
		First(&order)

	if result.Error != nil {
		return nil, result.Error
	}

	return &order, nil
}

func (r *orderRepo) CreateOrder(ctx context.Context, order *model.Order) (*model.Order, error) {
	result := r.db.Create(order)
	return order, result.Error
}
