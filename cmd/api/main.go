package main

import (
	"context"
	"log"

	"github.com/kpiasecki/wms/internal/repository/postgres"
)

func main() {
	db, err := postgres.NewDatabase()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close(context.Background())

	log.Println("Database connected")
}
