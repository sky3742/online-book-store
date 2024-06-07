package internal

import (
	repoAccount "online-book-store/internal/repository/account"
	repoBook "online-book-store/internal/repository/book"
	repoOrder "online-book-store/internal/repository/order"
	"online-book-store/internal/service"
	servAccount "online-book-store/internal/service/account"
	servBook "online-book-store/internal/service/book"
	servOrder "online-book-store/internal/service/order"
	"online-book-store/internal/utils"
)

type Service struct {
	Account service.Account
	Book    service.Book
	Order   service.Order
}

func InitServices() *Service {
	db := utils.InitDB(utils.Config.DatabaseUrl)

	// repository
	accountRepo := repoAccount.NewAccountRepo(db)
	bookRepo := repoBook.NewBookRepo(db)
	orderRepo := repoOrder.NewOrderRepo(db)

	// service
	accountService := servAccount.NewAccountService(servAccount.AccountServiceConfig{
		AccountRepo: accountRepo,
	})
	bookService := servBook.NewBookService(servBook.BookServiceConfig{
		BookRepo: bookRepo,
	})
	orderService := servOrder.NewOrderService(servOrder.OrderServiceConfig{
		BookRepo:  bookRepo,
		OrderRepo: orderRepo,
	})

	return &Service{
		Account: accountService,
		Book:    bookService,
		Order:   orderService,
	}
}
