package service

import (
	"cushon/internal/model"
	"cushon/internal/repository"
)

// Customer defines the interface for customer operations
type Customer interface {
	Create(customer *model.Customer) error
	Get(id int) (*model.Customer, error)
	Update(customer *model.Customer) error
}

// defaultCustomerService is a concrete implementation of CustomerService.
type defaultCustomerService struct {
	repo repository.CustomerRepository
}

// NewDefaultCustomerService creates a new default user service.
func NewDefaultCustomerService(repo repository.CustomerRepository) *defaultCustomerService {
	return &defaultCustomerService{repo: repo}
}
