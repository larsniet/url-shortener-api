package url

// Repository defines the interface for URL data access
type Repository interface {
	GetByID(id string) (URL, error)
	GetBySlug(slug string) (string, error)
	Save(originalURL string) (string, string, error)
	Delete(id string) error
}
