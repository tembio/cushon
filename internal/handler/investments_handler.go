package handler

import (
	"cushon/internal/service"
	"net/http"
)

// InvestmentHandler handles investment-related HTTP requests
type InvestmentHandler struct {
	investmentService service.Investment
}

// NewInvestmentHandler creates a new investment handler
func NewInvestmentHandler(investmentService service.Investment) *InvestmentHandler {
	return &InvestmentHandler{
		investmentService: investmentService,
	}
}

// Create handles investment creation
func (h *InvestmentHandler) Create(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not implemented", http.StatusNotImplemented)
}

// Get handles retrieving an investment
func (h *InvestmentHandler) Get(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not implemented", http.StatusNotImplemented)
}

// Update handles investment updates
func (h *InvestmentHandler) Update(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not implemented", http.StatusNotImplemented)
}

// Delete handles investment deletion
func (h *InvestmentHandler) Delete(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not implemented", http.StatusNotImplemented)
}
