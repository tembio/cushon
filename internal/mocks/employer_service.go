package mocks

import (
	"cushon/internal/model"
)

// EmployerService is a mock implementation of service.Employer
type EmployerService struct {
	MockEmployer  *model.Employer
	MockEmployers []*model.Employer
	MockErr       error
}

// NewEmployer implements service.Employer
func (m *EmployerService) NewEmployer(name string) (*model.Employer, error) {
	if m.MockErr != nil {
		return nil, m.MockErr
	}
	return m.MockEmployer, nil
}

// GetAllEmployers implements service.Employer
func (m *EmployerService) GetAllEmployers() ([]*model.Employer, error) {
	if m.MockErr != nil {
		return nil, m.MockErr
	}
	return m.MockEmployers, nil
}
