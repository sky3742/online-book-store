package account

import (
	"context"
	"online-book-store/internal/constant"
	"online-book-store/internal/model"
	"online-book-store/internal/utils"

	"github.com/gofiber/fiber/v2/log"
)

func (s *accountService) RegisterAccount(ctx context.Context, account *model.Account) (*model.Account, *model.AppError) {
	account.HashedPassword = utils.GenerateHashPassword(account.Email, account.Password)

	account, err := s.AccountRepo.RegisterAccount(ctx, account)
	if err != nil {
		log.Errorf("error register account: %v", err)
		return nil, constant.ErrRegisteredAccount
	}

	return account, nil
}

func (s *accountService) GetAccount(ctx context.Context, account *model.Account) (*model.Account, *model.AppError) {
	account.HashedPassword = utils.GenerateHashPassword(account.Email, account.Password)

	account, err := s.AccountRepo.GetAccount(ctx, account)
	if err != nil {
		log.Errorf("error get account: %v", err)
		return nil, constant.ErrAccountNotFound
	}

	return account, nil
}
