package files

import (
	"online-book-store/internal/model"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

func CreateBookTable() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "2024060602_create_book_table",

		Migrate: func(db *gorm.DB) error {
			type Book struct {
				model.ModelBase
				Title           string
				Description     string
				Author          string
				PublicationYear int
				Genre           datatypes.JSONSlice[string]
				CoverImage      string
				Price           float64
			}

			return db.AutoMigrate(&Book{})
		},

		Rollback: func(db *gorm.DB) error {
			return db.Migrator().DropTable("books")
		},
	}
}
