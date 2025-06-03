package handler

import (
	"cushon/internal/service"
	"net/http"
)

// EmployerHandler handles employer-related HTTP requests
type EmployerHandler struct {
	employerService service.Employer
}

// NewEmployerHandler creates a new employer handler
func NewEmployerHandler(employerService service.Employer) *EmployerHandler {
	return &EmployerHandler{
		employerService: employerService,
	}
}

// Create handles employer creation
func (h *EmployerHandler) Create(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not implemented", http.StatusNotImplemented)
}

// Get handles retrieving an employer
func (h *EmployerHandler) Get(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not implemented", http.StatusNotImplemented)
}

// Update handles employer updates
func (h *EmployerHandler) Update(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not implemented", http.StatusNotImplemented)
}

// Delete handles employer deletion
func (h *EmployerHandler) Delete(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not implemented", http.StatusNotImplemented)
}
