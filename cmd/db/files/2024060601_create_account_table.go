package files

import (
	"online-book-store/internal/model"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func CreateAccountTable() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "2024060601_create_account_table",

		Migrate: func(db *gorm.DB) error {
			type Account struct {
				model.ModelBase
				Email          string `gorm:"unique"`
				HashedPassword string
			}

			return db.AutoMigrate(&Account{})
		},

		Rollback: func(db *gorm.DB) error {
			return db.Migrator().DropTable("accounts")
		},
	}
}
