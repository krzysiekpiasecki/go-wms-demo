package main

import (
	"context"
	"log"

	"github.com/kpiasecki/wms/internal/config"
	"github.com/kpiasecki/wms/internal/repository/postgres"
)

func main() {
	cfg := config.Load()

	db, err := postgres.NewDatabase(cfg)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close(context.Background())

	var version string

	err = db.QueryRow(
		context.Background(),
		"SELECT version()",
	).Scan(&version)

	if err != nil {
		log.Fatal(err)
	}

	log.Println(version)

	log.Println("Database connected")
}
