package main

import (
	"log"
	"online-book-store/cmd/db"
)

func main() {
	err := db.Migrations.RollbackLast()
	if err != nil {
		log.Fatalf("failed to rollback: %s", err)
	}
}
