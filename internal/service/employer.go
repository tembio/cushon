package service

import (
	"cushon/internal/model"
	"cushon/internal/repository"
)

// Employer defines the interface for employer operations
type Employer interface {
	NewEmployer(name string) (*model.Employer, error)
}

// defaultEmployerService is a concrete implementation of Employer
type defaultEmployerService struct {
	repo repository.EmployerRepository
}

// NewDefaultEmployerService creates a new default employer service
func NewDefaultEmployerService(repo repository.EmployerRepository) *defaultEmployerService {
	return &defaultEmployerService{repo: repo}
}

// Create creates a new employer
func (s *defaultEmployerService) NewEmployer(name string) (*model.Employer, error) {
	return s.repo.CreateEmployer(name)
}
