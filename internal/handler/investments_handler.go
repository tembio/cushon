package handler

import (
	"cushon/internal/model"
	"cushon/internal/service"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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
	var createRequest model.InvestmentCreate
	if err := json.NewDecoder(r.Body).Decode(&createRequest); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate request
	if createRequest.ClientID == 0 {
		http.Error(w, "Client ID is required", http.StatusBadRequest)
		return
	}
	if createRequest.FundID == 0 {
		http.Error(w, "Fund ID is required", http.StatusBadRequest)
		return
	}
	if createRequest.Amount <= 0 {
		http.Error(w, "Amount must be greater than 0", http.StatusBadRequest)
		return
	}

	investment, err := h.investmentService.NewInvestment(createRequest.ClientID, createRequest.FundID, float32(createRequest.Amount))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := model.InvestmentResponse{
		ID:       investment.ID,
		ClientID: investment.ClientID,
		FundID:   investment.FundID,
		Amount:   float64(investment.Amount),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

// Get handles retrieving an investment
func (h *InvestmentHandler) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid investment ID", http.StatusBadRequest)
		return
	}

	investment, err := h.investmentService.GetInvestment(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	response := model.InvestmentResponse{
		ID:       investment.ID,
		ClientID: investment.ClientID,
		FundID:   investment.FundID,
		Amount:   float64(investment.Amount),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
