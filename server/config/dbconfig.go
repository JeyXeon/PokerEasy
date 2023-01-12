package config

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
	"os"
)

const (
	dbURL = "postgres://postgres:123@localhost:5432/poker_easy"
)

func GetDbConnection() *pgxpool.Pool {
	db, err := pgxpool.New(context.Background(), dbURL)
	if err != nil {
		logrus.WithError(err).Info("Unable to connect to database: %s\n", err.Error())
		os.Exit(1)
	}

	return db
}
