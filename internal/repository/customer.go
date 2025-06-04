package repository

import (
	"cushon/internal/model"
	"errors"
	"time"
)

// CustomerRepository defines the contract for storing and retrieving user data.
type CustomerRepository interface {
	CreateCustomer(customerName string, employerID *uint) (*model.Customer, error)
}

// InMemoryCustomerRepository is a simple in-memory implementation of CustomerRepository for demonstration.
type InMemoryCustomerRepository struct {
	customers map[uint]*model.Customer
	nextID    uint
}

// NewInMemoryCustomerRepository creates a new in-memory customer repository.
func NewInMemoryCustomerRepository() *InMemoryCustomerRepository {
	return &InMemoryCustomerRepository{
		customers: make(map[uint]*model.Customer),
		nextID:    1,
	}
}

// CreateCustomer creates a new customer
func (r *InMemoryCustomerRepository) CreateCustomer(customerName string, employerID *uint) (*model.Customer, error) {
	if customerName == "" {
		return nil, errors.New("customer name cannot be empty")
	}

	now := time.Now()
	customer := &model.Customer{
		ID:         r.nextID,
		Name:       customerName,
		EmployerID: employerID,
		CreatedAt:  now,
		UpdatedAt:  now,
	}

	r.customers[customer.ID] = customer
	r.nextID++

	return customer, nil
}
