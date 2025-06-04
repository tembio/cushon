package mocks

import (
	"cushon/internal/model"
)

// EmployerRepository is a mock implementation of repository.EmployerRepository
type EmployerRepository struct {
	CreateEmployerErr error
	MockEmployer      *model.Employer
}

// CreateEmployer implements repository.EmployerRepository
func (m *EmployerRepository) CreateEmployer(name string) (*model.Employer, error) {
	if m.CreateEmployerErr != nil {
		return nil, m.CreateEmployerErr
	}
	return m.MockEmployer, nil
}
