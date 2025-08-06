package url

import (
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

// GetShortURLHandler godoc
// @Summary Get a shortened URL
// @Tags URL Management
// @Description Gets a shortened URL by ID
// @Accept json
// @Produce json
// @Param id path string true "Short URL ID"
// @Success 200 {object} GetShortURLResponse
// @Failure 404 {object} map[string]string
// @Router /urls/{id} [get]
func (h *Handler) GetShortURLHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	h.service.GetShortURL(w, r, id)
}

// CreateShortURLHandler godoc
// @Summary Shorten a URL
// @Tags URL Management
// @Description Generates a shortened URL
// @Accept json
// @Produce json
// @Param request body CreateShortURLRequest true "Original URL"
// @Success 200 {object} CreateShortURLResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /urls [post]
func (h *Handler) CreateShortURLHandler(w http.ResponseWriter, r *http.Request) {
	h.service.CreateShortURL(w, r)
}

// DeleteShortURLHandler godoc
// @Summary Delete a shortened URL
// @Tags URL Management
// @Description Deletes a shortened URL by ID
// @Accept json
// @Produce json
// @Param request body DeleteShortURLRequest true "URL ID"
// @Success 200 {object} DeleteShortURLResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /urls [delete]
func (h *Handler) DeleteShortURLHandler(w http.ResponseWriter, r *http.Request) {
	h.service.DeleteShortURL(w, r)
}

// RedirectHandler godoc
// @Summary Redirect to original URL
// @Tags URL Redirect Route
// @Description Redirects to the original URL by slug
// @Accept json
// @Produce json
// @Param slug path string true "Slug"
// @Success 302 {string} string "Redirects to original URL"
// @Failure 404 {object} map[string]string
// @Router /{slug} [get]
func (h *Handler) RedirectHandler(w http.ResponseWriter, r *http.Request) {
	slug := strings.TrimPrefix(r.URL.Path, "/")
	h.service.RedirectURL(w, r, slug)
}
