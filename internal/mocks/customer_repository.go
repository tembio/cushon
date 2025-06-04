package mocks

import (
	"cushon/internal/model"
)

// CustomerRepository is a mock implementation of repository.CustomerRepository
type CustomerRepository struct {
	CreateCustomerErr error
	MockCustomer      *model.Customer
}

// CreateCustomer implements repository.CustomerRepository
func (m *CustomerRepository) CreateCustomer(customerName string, employerID *uint) (*model.Customer, error) {
	if m.CreateCustomerErr != nil {
		return nil, m.CreateCustomerErr
	}
	return m.MockCustomer, nil
}
