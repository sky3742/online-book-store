package service

import (
	"context"
	"online-book-store/internal/model"
)

type Account interface {
	CreateAccount(ctx context.Context, account *model.Account) (*model.Account, error)
	GetAccount(ctx context.Context, account *model.Account) (*model.Account, error)
}

type Book interface {
	GetBooks(ctx context.Context) ([]model.Book, error)
	GetBook(ctx context.Context, id string) (*model.Book, error)
}

type Order interface {
	GetOrders(ctx context.Context, userID string) ([]model.Order, error)
	GetOrder(ctx context.Context, userID, orderNo string) (*model.Order, error)
	CreateOrder(ctx context.Context, userID string, order *model.Order) (*model.Order, error)
}
