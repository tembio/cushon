package repository

import (
	"cushon/internal/model"
	"errors"
)

// EmployerRepository defines the contract for storing and retrieving employer data
type EmployerRepository interface {
	CreateEmployer(name string) (*model.Employer, error)
}

// InMemoryEmployerRepository is a simple in-memory implementation of EmployerRepository
type InMemoryEmployerRepository struct {
	employers map[uint]*model.Employer
	nextID    uint
}

// NewInMemoryEmployerRepository creates a new in-memory employer repository
func NewInMemoryEmployerRepository() *InMemoryEmployerRepository {
	return &InMemoryEmployerRepository{
		employers: make(map[uint]*model.Employer),
		nextID:    1,
	}
}

// Create creates a new employer
func (r *InMemoryEmployerRepository) CreateEmployer(name string) (*model.Employer, error) {
	if name == "" {
		return nil, errors.New("employer name cannot be empty")
	}

	employer := &model.Employer{
		ID:   r.nextID,
		Name: name,
	}

	r.employers[employer.ID] = employer
	r.nextID++

	return employer, nil
}
