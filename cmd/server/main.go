package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"url-shortener/internal/db"
	"url-shortener/internal/handler"
	"url-shortener/pkg/logger"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	logger.InitLogger()
	logger.Info("Starting URL shortener server...")

	connectionString := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DB"),
	)

	if err := db.InitDB(connectionString); err != nil {
		logger.Error("Failed to connect to database: %v", err)
		os.Exit(1)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/health-check", handler.HealthCheckHandler)
	mux.HandleFunc("/urls", handler.CreateShortURLHandler)
	mux.HandleFunc("/", handler.RedirectHandler)

	port := ":" + os.Getenv("APP_PORT")
	logger.Info("Listening on %s", port)
	if err := http.ListenAndServe(port, mux); err != nil {
		logger.Error("Server error: %v", err)
	}
}
