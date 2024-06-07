package order

import (
	"context"
	"errors"
	"online-book-store/internal/model"
	"online-book-store/internal/repository"
	repoMock "online-book-store/internal/repository/mock"
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

func Test_orderService_GetOrders(t *testing.T) {
	ctx := context.Background()
	userID := "123"

	type fields struct {
		BookRepo  *repoMock.BookProvider
		OrderRepo *repoMock.OrderProvider
	}
	type args struct {
		ctx    context.Context
		userID string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []model.Order
		wantErr bool
	}{
		{
			name: "Should get orders",
			fields: fields{
				BookRepo: func() *repoMock.BookProvider {
					mock := repoMock.BookProvider{}
					return &mock
				}(),
				OrderRepo: func() *repoMock.OrderProvider {
					mock := repoMock.OrderProvider{}
					mock.On(
						"GetOrders",
						ctx,
						userID,
					).Return(
						[]model.Order{
							{},
						},
						nil,
					)
					return &mock
				}(),
			},
			args: args{
				ctx:    ctx,
				userID: userID,
			},
			want: []model.Order{
				{},
			},
			wantErr: false,
		},
		{
			name: "Should get empty order",
			fields: fields{
				BookRepo: func() *repoMock.BookProvider {
					mock := repoMock.BookProvider{}
					return &mock
				}(),
				OrderRepo: func() *repoMock.OrderProvider {
					mock := repoMock.OrderProvider{}
					mock.On(
						"GetOrders",
						ctx,
						userID,
					).Return(
						[]model.Order{},
						nil,
					)
					return &mock
				}(),
			},
			args: args{
				ctx:    ctx,
				userID: userID,
			},
			want:    []model.Order{},
			wantErr: false,
		},
		{
			name: "Should fail to get orders",
			fields: fields{
				BookRepo: func() *repoMock.BookProvider {
					mock := repoMock.BookProvider{}
					return &mock
				}(),
				OrderRepo: func() *repoMock.OrderProvider {
					mock := repoMock.OrderProvider{}
					mock.On(
						"GetOrders",
						ctx,
						userID,
					).Return(
						[]model.Order{},
						errors.New("error"),
					)
					return &mock
				}(),
			},
			args: args{
				ctx:    ctx,
				userID: userID,
			},
			want:    []model.Order{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &orderService{
				BookRepo:  tt.fields.BookRepo,
				OrderRepo: tt.fields.OrderRepo,
			}
			got, err := s.GetOrders(tt.args.ctx, tt.args.userID)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func Test_orderService_GetOrder(t *testing.T) {
	ctx := context.Background()
	userID := "123"
	orderNo := "456"

	type fields struct {
		BookRepo  repository.BookProvider
		OrderRepo repository.OrderProvider
	}
	type args struct {
		ctx     context.Context
		userID  string
		orderNo string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.Order
		wantErr bool
	}{
		{
			name: "Should get order",
			fields: fields{
				BookRepo: func() repository.BookProvider {
					mock := repoMock.BookProvider{}
					return &mock
				}(),
				OrderRepo: func() repository.OrderProvider {
					mock := repoMock.OrderProvider{}
					mock.On(
						"GetOrder",
						ctx,
						userID,
						orderNo,
					).Return(
						&model.Order{},
						nil,
					)
					return &mock
				}(),
			},
			args: args{
				ctx:     ctx,
				userID:  userID,
				orderNo: orderNo,
			},
			want:    &model.Order{},
			wantErr: false,
		},
		{
			name: "Should fail to get order",
			fields: fields{
				BookRepo: func() repository.BookProvider {
					mock := repoMock.BookProvider{}
					return &mock
				}(),
				OrderRepo: func() repository.OrderProvider {
					mock := repoMock.OrderProvider{}
					mock.On(
						"GetOrder",
						ctx,
						userID,
						orderNo,
					).Return(
						nil,
						errors.New("error"),
					)
					return &mock
				}(),
			},
			args: args{
				ctx:     ctx,
				userID:  userID,
				orderNo: orderNo,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &orderService{
				BookRepo:  tt.fields.BookRepo,
				OrderRepo: tt.fields.OrderRepo,
			}
			got, err := s.GetOrder(tt.args.ctx, tt.args.userID, tt.args.orderNo)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func Test_orderService_CreateOrder(t *testing.T) {
	ctx := context.Background()
	userID := uuid.NewV4()
	order := &model.Order{
		Items: []model.OrderItem{
			{
				BookID:   uuid.NewV4(),
				Quantity: 1,
			},
		},
	}
	bookPrice := float64(100)

	type fields struct {
		BookRepo  repository.BookProvider
		OrderRepo repository.OrderProvider
	}
	type args struct {
		ctx    context.Context
		userID string
		order  *model.Order
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.Order
		wantErr bool
	}{
		{
			name: "Should create order",
			fields: fields{
				BookRepo: func() repository.BookProvider {
					mock := repoMock.BookProvider{}
					mock.On(
						"GetBook",
						ctx,
						order.Items[0].BookID.String(),
					).Return(
						&model.Book{
							Price: bookPrice,
						},
						nil,
					)
					return &mock
				}(),
				OrderRepo: func() repository.OrderProvider {
					mock := repoMock.OrderProvider{}
					mock.On(
						"CreateOrder",
						ctx,
						&model.Order{
							UserID: userID,
							Status: "pending",
							Items: []model.OrderItem{
								{
									BookID:   order.Items[0].BookID,
									Quantity: order.Items[0].Quantity,
									Price:    bookPrice,
								},
							},
						},
					).Return(
						&model.Order{
							UserID: userID,
							Status: "pending",
							Items: []model.OrderItem{
								{
									BookID:   order.Items[0].BookID,
									Quantity: order.Items[0].Quantity,
									Price:    bookPrice,
								},
							},
						},
						nil,
					)
					return &mock
				}(),
			},
			args: args{
				ctx:    ctx,
				userID: userID.String(),
				order:  order,
			},
			want: &model.Order{
				UserID: userID,
				Status: "pending",
				Items: []model.OrderItem{
					{
						BookID:   order.Items[0].BookID,
						Quantity: order.Items[0].Quantity,
						Price:    bookPrice,
					},
				},
				TotalPrice: 100,
			},
			wantErr: false,
		},
		{
			name: "Should fail to create order",
			fields: fields{
				BookRepo: func() repository.BookProvider {
					mock := repoMock.BookProvider{}
					mock.On(
						"GetBook",
						ctx,
						order.Items[0].BookID.String(),
					).Return(
						&model.Book{
							Price: bookPrice,
						},
						nil,
					)
					return &mock
				}(),
				OrderRepo: func() repository.OrderProvider {
					mock := repoMock.OrderProvider{}
					mock.On(
						"CreateOrder",
						ctx,
						&model.Order{
							UserID: userID,
							Status: "pending",
							Items: []model.OrderItem{
								{
									BookID:   order.Items[0].BookID,
									Quantity: order.Items[0].Quantity,
									Price:    bookPrice,
								},
							},
						},
					).Return(
						nil,
						errors.New("error"),
					)
					return &mock
				}(),
			},
			args: args{
				ctx:    ctx,
				userID: userID.String(),
				order:  order,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Should fail to create order when book not found",
			fields: fields{
				BookRepo: func() repository.BookProvider {
					mock := repoMock.BookProvider{}
					mock.On(
						"GetBook",
						ctx,
						order.Items[0].BookID.String(),
					).Return(
						nil,
						errors.New("error"),
					)
					return &mock
				}(),
				OrderRepo: func() repository.OrderProvider {
					mock := repoMock.OrderProvider{}
					return &mock
				}(),
			},
			args: args{
				ctx:    ctx,
				userID: userID.String(),
				order:  order,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &orderService{
				BookRepo:  tt.fields.BookRepo,
				OrderRepo: tt.fields.OrderRepo,
			}
			got, err := s.CreateOrder(tt.args.ctx, tt.args.userID, tt.args.order)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}
