package handler

import (
	"cushon/internal/model"
	"cushon/internal/service"
	"encoding/json"
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
	var createRequest model.EmployerCreate
	if err := json.NewDecoder(r.Body).Decode(&createRequest); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	employer, err := h.employerService.NewEmployer(createRequest.Name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := model.EmployerResponse{
		ID:   employer.ID,
		Name: employer.Name,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}
