package db

import (
	"database/sql"
	"url-shortener/pkg/logger"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB(connectionString string) error {
	var err error
	DB, err = sql.Open("postgres", connectionString)

	if err != nil {
		logger.Error("failed to open db: %v", err)
		return err
	}

	if err := DB.Ping(); err != nil {
		logger.Error("failed to ping db: %v", err)
		return err
	}

	logger.Info("Database connection established")

	return err
}

func GetDB() *sql.DB {
	return DB
}
