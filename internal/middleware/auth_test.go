package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"cushon/internal/repository"
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
	valid_key := "valid-key"
	repo := repository.NewInMemoryAPIKeyRepository()
	repo.AddKey(valid_key)

	tests := []struct {
		name           string
		apiKey         string
		expectedStatus int
		expectedBody   string
		shouldCallNext bool
	}{
		{
			name:           "Missing API Key",
			apiKey:         "",
			expectedStatus: http.StatusUnauthorized,
			expectedBody:   "API key required\n",
			shouldCallNext: false,
		},
		{
			name:           "Invalid API Key",
			apiKey:         "invalid-key",
			expectedStatus: http.StatusForbidden,
			expectedBody:   "Invalid API key\n",
			shouldCallNext: false,
		},
		{
			name:           "Valid API Key",
			apiKey:         valid_key,
			expectedStatus: http.StatusOK,
			expectedBody:   "",
			shouldCallNext: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", "/", nil)
			if tt.apiKey != "" {
				req.Header.Set("X-API-Key", tt.apiKey)
			}

			rr := httptest.NewRecorder()
			handler := &mockHandler{}
			middleware := AuthMiddleware(repo, handler)
			middleware.ServeHTTP(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v",
					rr.Code, tt.expectedStatus)
			}

			// Check response body
			body := rr.Body.String()
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
