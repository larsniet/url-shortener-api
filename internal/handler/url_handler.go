package handler

import (
	"fmt"
	"net/http"
	"url-shortener/internal/services"

	"github.com/go-chi/chi/v5"
)

// HealthCheckHandler godoc
// @Summary Check if the server is running
// @Tags Health Check
// @Description Returns a simple message indicating the server is up and running
// @Accept json
// @Produce text/plain
// @Success 200 {string} string "Server is up and running!"
// @Router /health-check [get]
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Server is up and running!")
}

// GetShortURLHandler godoc
// @Summary Get a shortened URL
// @Tags URL Management
// @Description Gets a shortened URL by ID
// @Accept json
// @Produce json
// @Param id path string true "Short URL ID"
// @Success 200 {object} services.GetShortURLResponse
// @Failure 404 {object} map[string]string
// @Router /urls/{id} [get]
func GetShortURLHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	services.GetShortURL(w, r, id)
}

// CreateShortURLHandler godoc
// @Summary Shorten a URL
// @Tags URL Management
// @Description Generates a shortened URL
// @Accept json
// @Produce json
// @Param request body services.CreateShortURLRequest true "Original URL"
// @Success 200 {object} services.CreateShortURLResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /urls [post]
func CreateShortURLHandler(w http.ResponseWriter, r *http.Request) {
	services.CreateShortURL(w, r)
}

// DeleteShortURLHandler godoc
// @Summary Delete a shortened URL
// @Tags URL Management
// @Description Deletes a shortened URL by ID
// @Accept json
// @Produce json
// @Param request body services.DeleteShortURLRequest true "URL ID"
// @Success 200 {object} services.DeleteShortURLResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /urls [delete]
func DeleteShortURLHandler(w http.ResponseWriter, r *http.Request) {
	services.DeleteShortURL(w, r)
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
func RedirectHandler(w http.ResponseWriter, r *http.Request) {
	services.RedirectURL(w, r)
}
