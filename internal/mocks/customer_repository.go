package mocks

import (
	"cushon/internal/model"
)

// CustomerRepository is a mock implementation of repository.CustomerRepository
type CustomerRepository struct {
	MockCustomer *model.Customer
	MockErr      error
}

// CreateCustomer implements repository.CustomerRepository
func (m *CustomerRepository) CreateCustomer(customerName string, employerID *uint) (*model.Customer, error) {
	if m.MockErr != nil {
		return nil, m.MockErr
	}
	return m.MockCustomer, nil
}
