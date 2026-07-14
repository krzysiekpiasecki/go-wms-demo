package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func NewDatabase() (*pgx.Conn, error) {
	connString := "host=localhost port=5432 user=wms_user password=wms_password dbname=wms sslmode=disable"

	conn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		return nil, fmt.Errorf("cannot connect to database: %w", err)
	}

	return conn, nil
}
