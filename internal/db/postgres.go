package db

import (
	"database/sql"
	"embed"
	"url-shortener/pkg/logger"

	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

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

	// Run migrations using Goose
	if err := runMigrations(); err != nil {
		logger.Error("failed to run migrations: %v", err)
		return err
	}

	logger.Info("Database migrations completed successfully")

	return nil
}

func runMigrations() error {
	// Set the embedded migrations source
	goose.SetBaseFS(embedMigrations)

	// Run migrations up to the latest version
	if err := goose.SetDialect("postgres"); err != nil {
		return err
	}

	if err := goose.Up(DB, "migrations"); err != nil {
		return err
	}

	return nil
}

func GetDB() *sql.DB {
	return DB
}
