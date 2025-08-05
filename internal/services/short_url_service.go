package services

import (
	"encoding/json"
	"net/http"
	"os"
	"url-shortener/internal/db"
	"url-shortener/pkg/logger"
	"url-shortener/pkg/utils"
)

type GetShortURLResponse struct {
	ID          string `json:"id"`
	OriginalURL string `json:"original_url"`
	ShortURL    string `json:"short_url"`
	CreatedAt   string `json:"created_at"`
}

type CreateShortURLRequest struct {
	OriginalURL string `json:"original_url"`
}

type CreateShortURLResponse struct {
	ID       string `json:"id"`
	ShortURL string `json:"short_url"`
}

type DeleteShortURLRequest struct {
	ID string `json:"id"`
}

type DeleteShortURLResponse struct {
	Message string `json:"message"`
}

func GetShortURL(w http.ResponseWriter, r *http.Request, id string) {
	url, err := db.GetURL(id)
	if err != nil {
		logger.Error("Failed to get URL: %v", err)
		utils.WriteError(w, http.StatusInternalServerError, "Failed to get URL")
		return
	}

	res := GetShortURLResponse{
		ID:          url.ID,
		OriginalURL: url.OriginalURL,
		ShortURL:    url.ShortSlug,
		CreatedAt:   url.CreatedAt,
	}

	utils.WriteJSON(w, http.StatusOK, res)
	logger.Info("Successfully got URL: %s", url.ID)
}

func CreateShortURL(w http.ResponseWriter, r *http.Request) {
	var req CreateShortURLRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.WriteError(w, http.StatusBadRequest, "Invalid JSON")
		return
	}

	if req.OriginalURL == "" {
		utils.WriteError(w, http.StatusBadRequest, "original_url is required")
		return
	}

	id, slug, err := db.SaveURL(req.OriginalURL)
	if err != nil {
		logger.Error("Failed to save URL: %v", err)
		utils.WriteError(w, http.StatusInternalServerError, "Failed to save URL")
		return
	}

	shortURL := os.Getenv("APP_HOST") + "/" + slug
	logger.Info("Successfully generated shortURL: %s", shortURL)

	res := CreateShortURLResponse{ID: id, ShortURL: shortURL}
	utils.WriteJSON(w, http.StatusOK, res)
}

func DeleteShortURL(w http.ResponseWriter, r *http.Request) {
	var req DeleteShortURLRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.WriteError(w, http.StatusBadRequest, "Invalid JSON")
		return
	}

	err := db.DeleteURL(req.ID)
	if err != nil {
		logger.Error("Failed to delete URL: %v", err)
		utils.WriteError(w, http.StatusInternalServerError, "Failed to delete URL")
		return
	}

	utils.WriteJSON(w, http.StatusOK, DeleteShortURLResponse{Message: "URL deleted successfully"})
	logger.Info("URL deleted successfully: %s", req.ID)
}
