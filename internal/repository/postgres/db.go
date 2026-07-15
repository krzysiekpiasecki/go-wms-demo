package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/kpiasecki/wms/internal/config"
)

func NewDatabase(cfg *config.Config) (*pgx.Conn, error) {

	connString := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
		cfg.DBSSLMode,
	)

	conn, err := pgx.Connect(
		context.Background(),
		connString,
	)

	if err != nil {
		return nil, fmt.Errorf("cannot connect to database: %w", err)
	}

	return conn, nil
}
