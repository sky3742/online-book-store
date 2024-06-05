package account

import (
	"context"
	"online-book-store/internal/model"
)

func (r *accountRepo) GetAccount(ctx context.Context, account *model.Account) (*model.Account, error) {
	result := r.db.Where(
		&model.Account{
			Email:          account.Email,
			HashedPassword: account.HashedPassword,
		},
	).First(&account)

	return account, result.Error
}

func (r *accountRepo) CreateAccount(ctx context.Context, account *model.Account) (*model.Account, error) {
	result := r.db.Create(&account)
	return account, result.Error
}
