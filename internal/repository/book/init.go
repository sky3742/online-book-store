package book

import (
	"online-book-store/internal/repository"

	"gorm.io/gorm"
)

type bookRepo struct {
	db *gorm.DB
}

func NewBookRepo(db *gorm.DB) repository.BookProvider {
	return &bookRepo{
		db: db,
	}
}
