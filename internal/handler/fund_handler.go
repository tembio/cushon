package handler

import (
	"net/http"

	"cushon/internal/service"
)

// FundHandler handles fund-related HTTP requests
type FundHandler struct {
	fundService service.Fund
}

// NewFundHandler creates a new fund handler
func NewFundHandler(fundService service.Fund) *FundHandler {
	return &FundHandler{
		fundService: fundService,
	}
}

// Create handles fund creation
func (h *FundHandler) Create(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not implemented", http.StatusNotImplemented)
}

// Get handles retrieving a fund
func (h *FundHandler) Get(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not implemented", http.StatusNotImplemented)
}
