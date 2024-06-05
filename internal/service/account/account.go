package account

import (
	"context"
	"online-book-store/internal/model"
	"online-book-store/internal/utils"
)

func (s *accountService) CreateAccount(ctx context.Context, account *model.Account) (*model.Account, error) {
	account.HashedPassword = utils.GenerateHashPassword(account.Email, account.Password)

	account, err := s.AccountRepo.CreateAccount(ctx, account)

	return account, err
}

func (s *accountService) LoginAccount(ctx context.Context, account *model.Account) (*model.Account, error) {
	account.HashedPassword = utils.GenerateHashPassword(account.Email, account.Password)

	account, err := s.AccountRepo.GetAccount(ctx, account)

	return account, err
}
