package mocks

import (
	"cushon/internal/model"
)

// CustomerService is a mock implementation of service.Customer
type CustomerService struct {
	MockCustomer *model.Customer
	MockErr      error
}

// NewRetailCustomer implements service.Customer
func (m *CustomerService) NewRetailCustomer(name string) (*model.Customer, error) {
	if m.MockErr != nil {
		return nil, m.MockErr
	}
	return m.MockCustomer, nil
}

// NewEmployedCustomer implements service.Customer
func (m *CustomerService) NewEmployedCustomer(name string, employerID uint) (*model.Customer, error) {
	if m.MockErr != nil {
		return nil, m.MockErr
	}
	return m.MockCustomer, nil
}
