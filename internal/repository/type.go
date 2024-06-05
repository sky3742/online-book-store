package repository

import (
	"context"
	"online-book-store/internal/model"
)

type AccountProvider interface {
	GetAccount(ctx context.Context, account *model.Account) (*model.Account, error)
	CreateAccount(ctx context.Context, account *model.Account) (*model.Account, error)
}

type BookProvider interface {
	GetBooks(ctx context.Context) ([]model.Book, error)
	GetBook(ctx context.Context, id string) (*model.Book, error)
}

type OrderProvider interface {
	GetOrders(ctx context.Context, userID string) ([]model.Order, error)
	GetOrder(ctx context.Context, userID, orderNo string) (*model.Order, error)
	CreateOrder(ctx context.Context, order *model.Order) (*model.Order, error)
}
