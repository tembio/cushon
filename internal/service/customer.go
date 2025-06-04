package service

import (
	"cushon/internal/model"
	"cushon/internal/repository"
)

// Customer defines the interface for customer operations
type Customer interface {
	NewRetailCustomer(name string) (*model.Customer, error)
	NewEmployedCustomer(name string, employerID uint) (*model.Customer, error)
}

// defaultCustomerService is a concrete implementation of CustomerService.
type defaultCustomerService struct {
	repo repository.CustomerRepository
}

// NewDefaultCustomerService creates a new default user service.
func NewDefaultCustomerService(repo repository.CustomerRepository) *defaultCustomerService {
	return &defaultCustomerService{repo: repo}
}

// NewRetailCustomer creates a new retail customer
func (s *defaultCustomerService) NewRetailCustomer(name string) (*model.Customer, error) {
	return s.repo.CreateCustomer(name, nil)
}

// NewEmployedCustomer creates a new employed customer
func (s *defaultCustomerService) NewEmployedCustomer(name string, employerID uint) (*model.Customer, error) {
	return s.repo.CreateCustomer(name, &employerID)
}
