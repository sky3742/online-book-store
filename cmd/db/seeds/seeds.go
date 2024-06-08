package main

import (
	"log"
	"online-book-store/cmd/db"
)

func main() {
	err := db.Seeds.Migrate()
	if err != nil {
		log.Fatalf("failed to migrate: %s", err)
	}
}
