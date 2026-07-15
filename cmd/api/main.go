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
}
