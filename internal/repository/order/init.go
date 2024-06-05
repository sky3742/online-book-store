package order

import (
	"online-book-store/internal/repository"

	"gorm.io/gorm"
)

type orderRepo struct {
	db *gorm.DB
}

func NewOrderRepo(db *gorm.DB) repository.OrderProvider {
	return &orderRepo{
		db: db,
	}
}
