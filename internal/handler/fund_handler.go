package handler

import (
	"cushon/internal/model"
	"cushon/internal/service"
	"encoding/json"
	"net/http"
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
	var createRequest model.FundCreate
	if err := json.NewDecoder(r.Body).Decode(&createRequest); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	fund, err := h.fundService.NewFund(createRequest.Name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := model.FundResponse{
		ID:   fund.ID,
		Name: fund.Name,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

// GetAll handles retrieving all funds
func (h *FundHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	funds, err := h.fundService.GetAllFunds()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := make([]model.FundResponse, len(funds))
	for i, fund := range funds {
		response[i] = model.FundResponse{
			ID:   fund.ID,
			Name: fund.Name,
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
