// @title URL Shortener API
// @version 1.0
// @description Simple API for shortening URLs.
// @host localhost:8080
// @BasePath /
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"url-shortener/internal/db"
	"url-shortener/internal/health"
	"url-shortener/internal/url"
	"url-shortener/pkg/logger"

	_ "url-shortener/docs"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	httpSwagger "github.com/swaggo/http-swagger"
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

	urlRepo := url.NewPostgresRepository(db.GetDB())
	urlService := url.NewService(urlRepo)
	urlHandler := url.NewHandler(urlService)

	r := chi.NewRouter()

	r.Get("/swagger/*", httpSwagger.WrapHandler)
	r.Get("/health-check", health.HealthCheckHandler)
	r.Route("/urls", func(r chi.Router) {
		r.Post("/", urlHandler.CreateShortURLHandler)
		r.Delete("/", urlHandler.DeleteShortURLHandler)
		r.Get("/{id}", urlHandler.GetShortURLHandler)
	})
	r.Get("/{slug}", urlHandler.RedirectHandler)

	port := ":" + os.Getenv("APP_PORT")
	logger.Info("Listening on %s", port)
	if err := http.ListenAndServe(port, r); err != nil {
		logger.Error("Server error: %v", err)
	}
}
