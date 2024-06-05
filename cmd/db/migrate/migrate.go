package main

import (
	"log"
	"online-book-store/cmd/db"
)

func main() {
	err := db.Migrations.Migrate()
	if err != nil {
		log.Fatalf("failed to migrate: %s", err)
	}
}
