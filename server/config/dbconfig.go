package config

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"os"
)

const (
	dbURL = "postgres://postgres:123@localhost:5432/poker_easy"
)

func GetDbConnection() *pgxpool.Pool {
	db, err := pgxpool.New(context.Background(), dbURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	return db
}
