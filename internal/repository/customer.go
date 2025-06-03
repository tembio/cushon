package repository

import "cushon/internal/model"

// CustomerRepository defines the contract for storing and retrieving user data.
type CustomerRepository interface {
	GetCustomerByID(id int) (*model.Customer, error)
}

// InMemoryCustomerRepository is a simple in-memory implementation of CustomerRepository for demonstration.
type InMemoryCustomerRepository struct {
	customers map[int]*model.Customer
}

// NewInMemoryCustomerRepository creates a new in-memory customer repository.
func NewInMemoryCustomerRepository() *InMemoryCustomerRepository {
	return &InMemoryCustomerRepository{
		customers: make(map[int]*model.Customer),
	}
}
