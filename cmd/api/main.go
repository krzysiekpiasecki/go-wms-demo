package main

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/kpiasecki/wms/internal/config"
)

func main() {
	conn, err := pgx.Connect(
		context.Background(),
		config.DatabaseURL,
	)
	if err != nil {
		panic(err)
	}
	defer conn.Close(context.Background())

	var dbName string

	err = conn.QueryRow(
		context.Background(),
		"SELECT current_database()",
	).Scan(&dbName)

	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to database:", dbName)
}
