package mocks

import (
	"cushon/internal/model"
)

// InvestmentRepository is a mock implementation of repository.InvestmentRepository for testing
type InvestmentRepository struct {
	CreateInvestment *model.Investment
	GetInvestment    *model.Investment

	CreateErr error
	GetErr    error
}

// Create implements the repository.InvestmentRepository interface
func (m *InvestmentRepository) Create(investment *model.Investment) error {
	if m.CreateErr != nil {
		return m.CreateErr
	}
	investment.ID = m.CreateInvestment.ID
	investment.CreatedAt = m.CreateInvestment.CreatedAt
	investment.UpdatedAt = m.CreateInvestment.UpdatedAt
	return nil
}

// GetByID implements the repository.InvestmentRepository interface
func (m *InvestmentRepository) GetByID(id uint) (*model.Investment, error) {
	if m.GetErr != nil {
		return nil, m.GetErr
	}
	return m.GetInvestment, nil
}
