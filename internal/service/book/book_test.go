package book

import (
	"context"
	"errors"
	"online-book-store/internal/model"
	"online-book-store/internal/repository"
	repoMock "online-book-store/internal/repository/mock"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_bookService_GetBooks(t *testing.T) {
	ctx := context.Background()

	type fields struct {
		BookRepo *repoMock.BookProvider
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []model.Book
		wantErr bool
	}{
		{
			name: "Should return books",
			fields: fields{
				BookRepo: func() *repoMock.BookProvider {
					mock := repoMock.BookProvider{}
					mock.On("GetBooks", ctx).Return(
						[]model.Book{
							{},
							{},
						},
						nil,
					)
					return &mock
				}(),
			},
			args: args{
				ctx: ctx,
			},
			want: []model.Book{
				{},
				{},
			},
			wantErr: false,
		},
		{
			name: "Should return empty array",
			fields: fields{
				BookRepo: func() *repoMock.BookProvider {
					mock := repoMock.BookProvider{}
					mock.On("GetBooks", ctx).Return(
						[]model.Book{},
						nil,
					)
					return &mock
				}(),
			},
			args: args{
				ctx: ctx,
			},
			want:    []model.Book{},
			wantErr: false,
		},
		{
			name: "Should fail",
			fields: fields{
				BookRepo: func() *repoMock.BookProvider {
					mock := repoMock.BookProvider{}
					mock.On("GetBooks", ctx).Return(
						[]model.Book{},
						errors.New("error"),
					)
					return &mock
				}(),
			},
			args: args{
				ctx: ctx,
			},
			want:    []model.Book{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &bookService{
				BookRepo: tt.fields.BookRepo,
			}
			got, err := s.GetBooks(tt.args.ctx)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func Test_bookService_GetBook(t *testing.T) {
	ctx := context.Background()
	id := "1"

	type fields struct {
		BookRepo repository.BookProvider
	}
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.Book
		wantErr bool
	}{
		{
			name: "Should return book",
			fields: fields{
				BookRepo: func() *repoMock.BookProvider {
					mock := repoMock.BookProvider{}
					mock.On("GetBook", ctx, id).Return(
						&model.Book{},
						nil,
					)
					return &mock
				}(),
			},
			args: args{
				ctx: ctx,
				id:  id,
			},
			want:    &model.Book{},
			wantErr: false,
		},
		{
			name: "Should fail",
			fields: fields{
				BookRepo: func() *repoMock.BookProvider {
					mock := repoMock.BookProvider{}
					mock.On("GetBook", ctx, id).Return(
						nil,
						errors.New("error"),
					)
					return &mock
				}(),
			},
			args: args{
				ctx: ctx,
				id: id,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &bookService{
				BookRepo: tt.fields.BookRepo,
			}
			got, err := s.GetBook(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("bookService.GetBook() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("bookService.GetBook() = %v, want %v", got, tt.want)
			}
		})
	}
}
