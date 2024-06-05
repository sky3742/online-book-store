package internal

import (
	"log"
	"online-book-store/internal/utils"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	db, err := gorm.Open(
		sqlite.Dialector{
			DriverName: "libsql",
			DSN:        utils.Config.DatabaseUrl,
		},
		&gorm.Config{},
	)

	if err != nil {
		log.Fatalf("failed to open db %s: %s", utils.Config.DatabaseUrl, err)
	}

	return db
}
