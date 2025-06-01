package handler

import (
	"cushon/internal/app/service"
	"net/http"
)

// ClientHandler handles client-related HTTP requests
type ClientHandler struct {
	clientService service.Client
}

// NewClientHandler creates a new client handler
func NewClientHandler(clientService service.Client) *ClientHandler {
	return &ClientHandler{
		clientService: clientService,
	}
}

// Create handles client creation
func (h *ClientHandler) Create(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not implemented", http.StatusNotImplemented)
}

// Get handles retrieving a client
func (h *ClientHandler) Get(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not implemented", http.StatusNotImplemented)
}

// Update handles client updates
func (h *ClientHandler) Update(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not implemented", http.StatusNotImplemented)
}

// Delete handles client deletion
func (h *ClientHandler) Delete(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not implemented", http.StatusNotImplemented)
}
