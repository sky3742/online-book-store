package utils

import (
	"log"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDB(dbUrl string) *gorm.DB {
	db, err := gorm.Open(
		sqlite.Dialector{
			DriverName: "libsql",
			DSN:        dbUrl,
		},
		&gorm.Config{},
	)

	if err != nil {
		log.Fatalf("failed to open db %s: %s", dbUrl, err)
	}

	return db
}
