package repository

// APIKeyRepository defines the interface for API key operations
type APIKeyRepository interface {
	// ValidateKey checks if an API key is valid
	ValidateKey(key string) bool
}

// InMemoryAPIKeyRepository implements APIKeyRepository using an in-memory store
type InMemoryAPIKeyRepository struct {
	keys map[string]bool
}

// NewInMemoryAPIKeyRepository creates a new instance of InMemoryAPIKeyRepository
func NewInMemoryAPIKeyRepository() *InMemoryAPIKeyRepository {
	return &InMemoryAPIKeyRepository{
		keys: make(map[string]bool),
	}
}

// ValidateKey implements the APIKeyRepository interface
func (r *InMemoryAPIKeyRepository) ValidateKey(key string) bool {
	return r.keys[key]
}

// AddKey adds a new API key to the repository
func (r *InMemoryAPIKeyRepository) AddKey(key string) {
	r.keys[key] = true
}
