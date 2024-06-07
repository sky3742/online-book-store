package account

import (
	"context"
	"errors"
	"online-book-store/internal/model"
	repoMock "online-book-store/internal/repository/mock"
	"online-book-store/internal/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_accountService_RegisterAccount(t *testing.T) {
	ctx := context.Background()

	email := "test@email.com"
	password := "123456"
	hashedPassword := utils.GenerateHashPassword(email, password)

	type fields struct {
		AccountRepo *repoMock.AccountProvider
	}
	type args struct {
		ctx     context.Context
		account *model.Account
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.Account
		wantErr bool
	}{
		{
			name: "Should create account",
			fields: fields{
				AccountRepo: func() *repoMock.AccountProvider {
					mock := repoMock.AccountProvider{}
					mock.On(
						"RegisterAccount",
						ctx,
						&model.Account{
							Email:          email,
							Password:       password,
							HashedPassword: hashedPassword,
						},
					).Return(
						&model.Account{
							Email:          email,
							HashedPassword: hashedPassword,
						},
						nil,
					)
					return &mock
				}(),
			},
			args: args{
				ctx: ctx,
				account: &model.Account{
					Email:    email,
					Password: password,
				},
			},
			want: &model.Account{
				Email:          email,
				HashedPassword: hashedPassword,
			},
			wantErr: false,
		},
		{
			name: "Should fail to create account",
			fields: fields{
				AccountRepo: func() *repoMock.AccountProvider {
					mock := repoMock.AccountProvider{}
					mock.On(
						"RegisterAccount",
						ctx,
						&model.Account{
							Email:          email,
							Password:       password,
							HashedPassword: hashedPassword,
						},
					).Return(
						nil,
						errors.New("error"),
					)
					return &mock
				}(),
			},
			args: args{
				ctx: ctx,
				account: &model.Account{
					Email:    email,
					Password: password,
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &accountService{
				AccountRepo: tt.fields.AccountRepo,
			}
			got, err := s.RegisterAccount(tt.args.ctx, tt.args.account)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func Test_accountService_GetAccount(t *testing.T) {
	ctx := context.Background()

	email := "test@email.com"
	password := "123456"
	hashedPassword := utils.GenerateHashPassword(email, password)

	type fields struct {
		AccountRepo *repoMock.AccountProvider
	}
	type args struct {
		ctx     context.Context
		account *model.Account
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.Account
		wantErr bool
	}{
		{
			name: "Should get account",
			fields: fields{
				AccountRepo: func() *repoMock.AccountProvider {
					mock := repoMock.AccountProvider{}
					mock.On(
						"GetAccount",
						ctx,
						&model.Account{
							Email:          email,
							Password:       password,
							HashedPassword: hashedPassword,
						},
					).Return(
						&model.Account{
							Email:          email,
							HashedPassword: hashedPassword,
						},
						nil,
					)
					return &mock
				}(),
			},
			args: args{
				ctx: ctx,
				account: &model.Account{
					Email:    email,
					Password: password,
				},
			},
			want: &model.Account{
				Email:          email,
				HashedPassword: hashedPassword,
			},
			wantErr: false,
		},
		{
			name: "Should fail to get account",
			fields: fields{
				AccountRepo: func() *repoMock.AccountProvider {
					mock := repoMock.AccountProvider{}
					mock.On(
						"GetAccount",
						ctx,
						&model.Account{
							Email:          email,
							Password:       password,
							HashedPassword: hashedPassword,
						},
					).Return(
						nil,
						errors.New("error"),
					)
					return &mock
				}(),
			},
			args: args{
				ctx: ctx,
				account: &model.Account{
					Email:    email,
					Password: password,
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &accountService{
				AccountRepo: tt.fields.AccountRepo,
			}
			got, err := s.GetAccount(tt.args.ctx, tt.args.account)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}
