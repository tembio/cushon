package mocks

import (
	"cushon/internal/model"
)

// InvestmentRepository is a mock implementation of repository.InvestmentRepository
type InvestmentRepository struct {
	CreateInvestmentErr error
	GetByIDErr          error
	MockInvestment      *model.Investment
}

// CreateInvestment implements repository.InvestmentRepository
func (m *InvestmentRepository) CreateInvestment(clientID, fundID uint, amount float32) (*model.Investment, error) {
	if m.CreateInvestmentErr != nil {
		return nil, m.CreateInvestmentErr
	}
	return m.MockInvestment, nil
}

// GetInvestmentByID implements repository.InvestmentRepository
func (m *InvestmentRepository) GetInvestmentByID(id uint) (*model.Investment, error) {
	if m.GetByIDErr != nil {
		return nil, m.GetByIDErr
	}
	return m.MockInvestment, nil
}
