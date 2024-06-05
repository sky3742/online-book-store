package internal

import (
	"online-book-store/internal/httpservice/account"
	"online-book-store/internal/httpservice/book"
	"online-book-store/internal/httpservice/order"
)

type HttpService struct {
	Account *account.AccountHandler
	Book    *book.BookHandler
	Order   *order.OrderHandler
}

func InitHttpService(service *Service) HttpService {
	return HttpService{
		Account: account.NewHandler(account.ConfigHandler{
			AccountService: service.Account,
		}),
		Book: book.NewHandler(book.ConfigHandler{
			BookService: service.Book,
		}),
		Order: order.NewHandler(order.ConfigHandler{
			OrderService: service.Order,
		}),
	}
}
