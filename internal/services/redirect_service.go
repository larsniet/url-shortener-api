package services

import (
	"net/http"
	"strings"
	"url-shortener/internal/db"
	"url-shortener/pkg/logger"
)

// RedirectURL godoc
// @Summary Redirect to original URL
// @Tags Redirect Route
// @Description Redirects to the original URL by slug
// @Accept json
// @Produce json
// @Param slug path string true "Slug"
// @Success 302 {string} string "Redirects to original URL"
// @Failure 404 {object} map[string]string
// @Router /{slug} [get]
func RedirectURL(w http.ResponseWriter, r *http.Request) {
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
