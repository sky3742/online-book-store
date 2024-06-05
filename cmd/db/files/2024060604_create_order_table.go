package files

import (
	"online-book-store/internal/model"

	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CreateOrderTable() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "2024060604_create_order_table",

		Migrate: func(db *gorm.DB) error {
			type OrderItem struct {
				model.ModelBase
				OrderID  uuid.UUID
				BookID   uuid.UUID
				Price    float64
				Quantity int
			}

			type Order struct {
				model.ModelBase
				UserID  uuid.UUID
				Status  string
				Items   []OrderItem
			}

			return db.AutoMigrate(&Order{}, &OrderItem{})
		},

		Rollback: func(db *gorm.DB) error {
			return db.Migrator().DropTable("orders", "order_items")
		},
	}
}
