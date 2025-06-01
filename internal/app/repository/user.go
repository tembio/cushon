package repository

import "cushon/internal/model"

// ClientRepository defines the contract for storing and retrieving user data.
// This interface allows us to swap out different storage implementations (e.g., in-memory, SQL...).
type ClientRepository interface {
	GetClientByID(id int) (*model.Client, error)
	// Add other methods like CreateUser, UpdateUser, DeleteUser
}

// inMemoryClientRepository is a simple in-memory implementation of ClientRepository for demonstration.
type inMemoryClientRepository struct {
	clients map[int]*model.Client
}

// NewInMemoryClientRepository creates a new in-memory client repository.
func NewInMemoryClientRepository() *inMemoryClientRepository {
	return &inMemoryClientRepository{
		clients: make(map[int]*model.Client),
	}
}
