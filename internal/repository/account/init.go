package account

import (
	"online-book-store/internal/repository"

	"gorm.io/gorm"
)

type accountRepo struct {
	db *gorm.DB
}

func NewAccountRepo(db *gorm.DB) repository.AccountProvider {
	return &accountRepo{
		db: db,
	}
}
