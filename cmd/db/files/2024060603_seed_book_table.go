package files

import (
	"encoding/json"
	"online-book-store/internal/model"
	"os"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

func SeedBookTable() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "2024060603_seed_book_table",

		Migrate: func(db *gorm.DB) error {
			type Book struct {
				model.ModelBase
				Title           string                      `json:"title"`
				Description     string                      `json:"description"`
				Author          string                      `json:"author"`
				PublicationYear int                         `json:"publication_year"`
				Genre           datatypes.JSONSlice[string] `json:"genre"`
				CoverImage      string                      `json:"cover_image"`
				Price           float64                     `json:"price"`
			}

			jsonData, err := os.ReadFile("./cmd/db/data/books.json")
			if err != nil {
				return err
			}

			var books []Book
			err = json.Unmarshal(jsonData, &books)
			if err != nil {
				return err
			}

			return db.Create(&books).Error
		},

		Rollback: func(db *gorm.DB) error {
			return nil
		},
	}
}
