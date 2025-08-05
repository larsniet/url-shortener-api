package db

import (
	"crypto/rand"
	"database/sql"
	"errors"
	"math/big"

	"github.com/lib/pq"
)

var DB *sql.DB

const (
	slugChars  = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	slugLength = 7
	maxRetries = 5
)

func InitDB(connectionString string) error {
	var err error
	DB, err = sql.Open("postgres", connectionString)
	return err
}

func GetOriginalURL(slug string) (string, error) {
	var original string
	err := DB.QueryRow("SELECT original_url FROM urls WHERE short_slug = $1", slug).Scan(&original)
	return original, err
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

func SaveURL(originalURL string) (string, string, error) {
	for i := 0; i < maxRetries; i++ {
		slug, err := generateSlug()
		if err != nil {
			// If an error occurred during slug generation, return error
			return "", "", err
		}

		var id string
		err = DB.QueryRow("INSERT INTO urls (original_url, short_slug) VALUES ($1, $2) RETURNING id", originalURL, slug).Scan(&id)
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
