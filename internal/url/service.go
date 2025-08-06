package url

import (
	"encoding/json"
	"net/http"
	"os"
	"url-shortener/pkg/logger"
	"url-shortener/pkg/utils"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

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

func (s *Service) GetShortURL(w http.ResponseWriter, r *http.Request, id string) {
	url, err := s.repo.GetByID(id)
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

func (s *Service) CreateShortURL(w http.ResponseWriter, r *http.Request) {
	var req CreateShortURLRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.WriteError(w, http.StatusBadRequest, "Invalid JSON")
		return
	}

	if req.OriginalURL == "" {
		utils.WriteError(w, http.StatusBadRequest, "original_url is required")
		return
	}

	id, slug, err := s.repo.Save(req.OriginalURL)
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

func (s *Service) DeleteShortURL(w http.ResponseWriter, r *http.Request) {
	var req DeleteShortURLRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.WriteError(w, http.StatusBadRequest, "Invalid JSON")
		return
	}

	err := s.repo.Delete(req.ID)
	if err != nil {
		logger.Error("Failed to delete URL: %v", err)
		utils.WriteError(w, http.StatusInternalServerError, "Failed to delete URL")
		return
	}

	utils.WriteJSON(w, http.StatusOK, DeleteShortURLResponse{Message: "URL deleted successfully"})
	logger.Info("URL deleted successfully: %s", req.ID)
}

func (s *Service) RedirectURL(w http.ResponseWriter, r *http.Request, slug string) {
	if slug == "" {
		http.NotFound(w, r)
		return
	}

	originalURL, err := s.repo.GetBySlug(slug)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	logger.Info("Redirecting to %s", originalURL)
	http.Redirect(w, r, originalURL, http.StatusFound)
}
