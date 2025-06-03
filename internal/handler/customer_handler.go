package handler

import (
	"cushon/internal/service"
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
	http.Error(w, "Not implemented", http.StatusNotImplemented)
}
