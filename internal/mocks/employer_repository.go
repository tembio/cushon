package mocks

import (
	"cushon/internal/model"
)

// EmployerRepository is a mock implementation of repository.EmployerRepository
type EmployerRepository struct {
	MockEmployer *model.Employer
	MockErr      error
}

// CreateEmployer implements repository.EmployerRepository
func (m *EmployerRepository) CreateEmployer(name string) (*model.Employer, error) {
	if m.MockErr != nil {
		return nil, m.MockErr
	}
	return m.MockEmployer, nil
}
