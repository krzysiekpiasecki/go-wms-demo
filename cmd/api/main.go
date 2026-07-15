package main

import (
	"context"

	"github.com/kpiasecki/wms/internal/config"
	"github.com/kpiasecki/wms/internal/logger"
	"github.com/kpiasecki/wms/internal/repository/postgres"
)

func main() {
	cfg := config.Load()

	db, err := postgres.NewDatabase(cfg)
	if err != nil {
		logger.Log.Fatal().
			Err(err).
			Msg("application failed")

	}

	defer db.Close(context.Background())

	logger.Log.Info().
		Str("component", "api").
		Msg("application started")

}
