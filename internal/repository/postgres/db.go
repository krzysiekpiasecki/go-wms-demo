package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kpiasecki/wms/internal/config"
)

func NewDatabase(cfg *config.Config) (*pgxpool.Pool, error) {

	connString := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
		cfg.DBSSLMode,
	)

	pool, err := pgxpool.New(
		context.Background(),
		connString,
	)

	if err != nil {
		return nil, fmt.Errorf(
			"cannot connect to database: %w",
			err,
		)
	}

	return pool, nil
}
