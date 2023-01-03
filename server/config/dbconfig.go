package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	dbURL = "postgres://postgres:123@localhost:5432/poker_easy"
)

func GetDbConnection() *gorm.DB {
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return db
}
