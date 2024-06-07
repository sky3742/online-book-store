package service

import (
	"context"
	"online-book-store/internal/model"
)

type Account interface {
	RegisterAccount(ctx context.Context, account *model.Account) (*model.Account, *model.AppError)
	GetAccount(ctx context.Context, account *model.Account) (*model.Account, *model.AppError)
}

type Book interface {
	GetBooks(ctx context.Context) ([]model.Book, *model.AppError)
	GetBook(ctx context.Context, id string) (*model.Book, *model.AppError)
}

type Order interface {
	GetOrders(ctx context.Context, userID string) ([]model.Order, *model.AppError)
	GetOrder(ctx context.Context, userID, orderNo string) (*model.Order, *model.AppError)
	CreateOrder(ctx context.Context, userID string, order *model.Order) (*model.Order, *model.AppError)
}
