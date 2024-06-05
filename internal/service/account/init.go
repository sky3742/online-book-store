package account

import (
	"online-book-store/internal/repository"
	"online-book-store/internal/service"
)

type accountService struct {
	AccountRepo repository.AccountProvider
}

type AccountServiceConfig struct {
	AccountRepo repository.AccountProvider
}

func NewAccountService(cfg AccountServiceConfig) service.Account {
	return &accountService{
		AccountRepo: cfg.AccountRepo,
	}
}
