package db

import (
	"online-book-store/cmd/db/files"
	"online-book-store/internal/utils"

	"github.com/go-gormigrate/gormigrate/v2"
)

var Migrations *gormigrate.Gormigrate

func init() {
	utils.InitEnv()

	db := utils.InitDB(utils.Config.DatabaseUrl)

	Migrations = gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		// add migration files
		files.CreateAccountTable(),
		files.CreateBookTable(),
		files.SeedBookTable(),
		files.CreateOrderTable(),
	})
}
