package url

import (
	"crypto/rand"
	"database/sql"
	"errors"
	"math/big"

	"github.com/lib/pq"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) *PostgresRepository {
	return &PostgresRepository{db: db}
}

const (
	slugChars  = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	slugLength = 7
	maxRetries = 5
)

func (r *PostgresRepository) GetByID(id string) (URL, error) {
	var url URL
	err := r.db.QueryRow("SELECT id, original_url, short_slug, created_at FROM urls WHERE id = $1", id).Scan(&url.ID, &url.OriginalURL, &url.ShortSlug, &url.CreatedAt)
	return url, err
}

func (r *PostgresRepository) GetBySlug(slug string) (string, error) {
	var original string
	err := r.db.QueryRow("SELECT original_url FROM urls WHERE short_slug = $1", slug).Scan(&original)
	return original, err
}

func (r *PostgresRepository) Save(originalURL string) (string, string, error) {
	for i := 0; i < maxRetries; i++ {
		slug, err := generateSlug()
		if err != nil {
			// If an error occurred during slug generation, return error
			return "", "", err
		}

		var id string
		err = r.db.QueryRow("INSERT INTO urls (original_url, short_slug) VALUES ($1, $2) RETURNING id", originalURL, slug).Scan(&id)
		if err == nil {
			// If insertion into DB was successful, return the slug
			return id, slug, nil
		}

		// Handle unique constraint violation (Postgres code 23505)
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
			continue // Slug collision → retry
		}

		// If something else happened, return the error
		return "", "", err
	}

	return "", "", errors.New("failed to generate unique slug after multiple retries")
}

func (r *PostgresRepository) Delete(id string) error {
	_, err := r.db.Exec("DELETE FROM urls WHERE id = $1", id)
	return err
}

func generateSlug() (string, error) {
	b := make([]byte, slugLength)
	for i := range b {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(slugChars))))
		if err != nil {
			return "", err
		}
		b[i] = slugChars[n.Int64()]
	}
	return string(b), nil
}
