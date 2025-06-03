package middleware

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"cushon/internal/app/repository"
)

// mockHandler is a simple http.Handler that records if it was called
type mockHandler struct {
	called bool
}

func (m *mockHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	m.called = true
	w.WriteHeader(http.StatusOK)
}

func TestAuthMiddleware(t *testing.T) {
	tests := []struct {
		name        string
		apiKey      string
		validAPIKey bool

		expectedStatus int
		expectedBody   string
		shouldCallNext bool
	}{
		{
			name:        "Missing API Key",
			apiKey:      "",
			validAPIKey: false,

			expectedStatus: http.StatusUnauthorized,
			expectedBody:   "API key required",
			shouldCallNext: false,
		},
		{
			name:        "Invalid API Key",
			apiKey:      "invalid-key",
			validAPIKey: false,

			expectedStatus: http.StatusForbidden,
			expectedBody:   "Invalid API key",
			shouldCallNext: false,
		},
		{
			name:        "Valid API Key",
			apiKey:      "valid-key",
			validAPIKey: true,

			expectedStatus: http.StatusOK,
			expectedBody:   "",
			shouldCallNext: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := repository.NewInMemoryAPIKeyRepository()
			if tt.validAPIKey {
				repo.AddKey(tt.apiKey)
			}

			req := httptest.NewRequest("GET", "/", nil)
			if tt.apiKey != "" {
				req.Header.Set("X-API-Key", tt.apiKey)
			}

			rr := httptest.NewRecorder()
			handler := &mockHandler{}
			middleware := AuthMiddleware(repo, handler)
			middleware.ServeHTTP(rr, req)

			if status := rr.Code; status != tt.expectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tt.expectedStatus)
			}

			// Check response body
			body := strings.TrimSpace(rr.Body.String())
			if body != tt.expectedBody {
				t.Errorf("handler returned unexpected body: got %v want %v",
					body, tt.expectedBody)
			}

			if handler.called != tt.shouldCallNext {
				t.Errorf("next handler called: got %v want %v",
					handler.called, tt.shouldCallNext)
			}
		})
	}
}
