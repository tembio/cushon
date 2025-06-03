package middleware

import (
	"net/http"

	"cushon/internal/app/repository"
)

// AuthMiddleware is a middleware that checks for a valid API key
func AuthMiddleware(apiKeyRepo repository.APIKeyRepository, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("X-API-Key")
		if apiKey == "" {
			http.Error(w, "API key required", http.StatusUnauthorized)
			return
		}

		if !apiKeyRepo.ValidateKey(apiKey) {
			http.Error(w, "Invalid API key", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// NewAuthMiddleware creates a mux.MiddlewareFunc for authentication
func NewAuthMiddleware(apiKeyRepo repository.APIKeyRepository) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return AuthMiddleware(apiKeyRepo, next)
	}
}
