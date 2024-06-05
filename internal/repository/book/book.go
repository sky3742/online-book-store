package book

import (
	"context"
	"online-book-store/internal/model"
)

func (r *bookRepo) GetBooks(ctx context.Context) ([]model.Book, error) {
	var books = []model.Book{}
	result := r.db.Find(&books)
	return books, result.Error
}

func (r *bookRepo) GetBook(ctx context.Context, id string) (*model.Book, error) {
	var book model.Book
	result := r.db.Where("id = ?", id).First(&book)
	if result.Error != nil {
		return nil, result.Error
	}
	return &book, nil
}
