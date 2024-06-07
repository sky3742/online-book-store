package book

import (
	"context"
	"online-book-store/internal/constant"
	"online-book-store/internal/model"

	"github.com/gofiber/fiber/v2/log"
)

func (s *bookService) GetBooks(ctx context.Context) ([]model.Book, *model.AppError) {
	books, err := s.BookRepo.GetBooks(ctx)
	if err != nil {
		log.Errorf("error get books: %v", err)
		return []model.Book{}, constant.ErrGetBooks
	}
	return books, nil
}

func (s *bookService) GetBook(ctx context.Context, id string) (*model.Book, *model.AppError) {
	book, err := s.BookRepo.GetBook(ctx, id)
	if err != nil {
		log.Errorf("error get book: %v", err)
		return nil, constant.ErrBookNotFound
	}
	return book, nil
}
