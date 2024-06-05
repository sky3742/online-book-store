package book

import (
	"online-book-store/internal/repository"
	"online-book-store/internal/service"
)

type bookService struct {
	BookRepo repository.BookProvider
}

type BookServiceConfig struct {
	BookRepo repository.BookProvider
}

func NewBookService(cfg BookServiceConfig) service.Book {
	return &bookService{
		BookRepo: cfg.BookRepo,
	}
}
