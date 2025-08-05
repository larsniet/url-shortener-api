package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"url-shortener/internal/db"
	"url-shortener/pkg/logger"
)

type createShortURLRequest struct {
	OriginalURL string `json:"original_url"`
}
type createShortURLResponse struct {
	ID       string `json:"id"`
	ShortURL string `json:"short_url"`
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Server is up and running!")
}

func CreateShortURLHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var req createShortURLRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Bad Request: Invalid JSON", http.StatusBadRequest)
		return
	}

	if req.OriginalURL == "" {
		http.Error(w, "Bad Request: original_url is required", http.StatusBadRequest)
		return
	}

	id, slug, err := db.SaveURL(req.OriginalURL)
	if err != nil {
		logger.Error("Failed to save URL: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	shortURL := os.Getenv("APP_HOST") + "/" + slug
	logger.Info("Succesfully generated shortURL: %v", shortURL)

	res := createShortURLResponse{ID: id, ShortURL: shortURL}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func RedirectHandler(w http.ResponseWriter, r *http.Request) {
	slug := strings.TrimPrefix(r.URL.Path, "/")
	if slug == "" {
		http.NotFound(w, r)
		return
	}

	originalURL, err := db.GetOriginalURL(slug)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	logger.Info("Redirecting to %s", originalURL)
	http.Redirect(w, r, originalURL, http.StatusFound)
}
