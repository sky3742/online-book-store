package model

import (
	uuid "github.com/satori/go.uuid"
)

type OrderItem struct {
	ModelBase
	OrderID  uuid.UUID `json:"-"`
	BookID   uuid.UUID `json:"book_id"`
	Price    float64   `json:"price"`
	Quantity int       `json:"quantity" validate:"min=1"`
}

type Order struct {
	ModelBase
	UserID     uuid.UUID   `json:"-"`
	Items      []OrderItem `json:"items"`
	Status     string      `json:"status"`
	TotalPrice float64     `gorm:"-" json:"total_price"`
}

func (o *Order) CalculateTotalPrice() {
	o.TotalPrice = 0

	for _, item := range o.Items {
		o.TotalPrice += item.Price * float64(item.Quantity)
	}
}
