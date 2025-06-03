package repository

import "cushon/internal/model"

// CustomerRepository defines the contract for storing and retrieving user data.
// This interface allows us to swap out different storage implementations (e.g., in-memory, SQL...).
type CustomerRepository interface {
	GetCustomerByID(id int) (*model.Customer, error)
	// Add other methods like CreateUser, UpdateUser, DeleteUser
}

// inMemoryCustomerRepository is a simple in-memory implementation of CustomerRepository for demonstration.
type inMemoryCustomerRepository struct {
	customers map[int]*model.Customer
}

// NewInMemoryCustomerRepository creates a new in-memory customer repository.
func NewInMemoryCustomerRepository() *inMemoryCustomerRepository {
	return &inMemoryCustomerRepository{
		customers: make(map[int]*model.Customer),
	}
}
