package book

import (
	"context"
	"online-book-store/internal/model"
)

func (s *bookService) GetBooks(ctx context.Context) ([]model.Book, error) {
	books, err := s.BookRepo.GetBooks(ctx)
	if err != nil {
		return []model.Book{}, err
	}
	return books, err
}

func (s *bookService) GetBook(ctx context.Context, id string) (*model.Book, error) {
	book, err := s.BookRepo.GetBook(ctx, id)
	if err != nil {
		return nil, err
	}
	return book, nil
}
