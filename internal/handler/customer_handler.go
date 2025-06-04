package handler

import (
	"cushon/internal/model"
	"cushon/internal/service"
	"encoding/json"
	"net/http"
)

// CustomerHandler handles customer-related HTTP requests
type CustomerHandler struct {
	customerService service.Customer
}

// NewCustomerHandler creates a new customer handler
func NewCustomerHandler(customerService service.Customer) *CustomerHandler {
	return &CustomerHandler{
		customerService: customerService,
	}
}

// Create handles customer creation
func (h *CustomerHandler) Create(w http.ResponseWriter, r *http.Request) {
	var createRequest model.CustomerCreate
	if err := json.NewDecoder(r.Body).Decode(&createRequest); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	var customer *model.Customer
	var err error

	if createRequest.EmployerID == nil {
		// Create retail customer
		customer, err = h.customerService.NewRetailCustomer(createRequest.Name)
	} else {
		// Create employed customer
		customer, err = h.customerService.NewEmployedCustomer(createRequest.Name, *createRequest.EmployerID)
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := model.CustomerResponse{
		ID:         customer.ID,
		Name:       customer.Name,
		EmployerID: customer.EmployerID,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}
