package mocks

import (
	"cushon/internal/model"
)

// InvestmentRepository is a mock implementation of the InvestmentRepository interface
type InvestmentRepository struct {
	MockInvestment  *model.Investment
	MockInvestments []*model.Investment
	MockErr         error
}

// CreateInvestment creates a new investment
func (m *InvestmentRepository) CreateInvestment(clientID, fundID uint, amount float32) (*model.Investment, error) {
	if m.MockErr != nil {
		return nil, m.MockErr
	}
	return m.MockInvestment, nil
}

// GetInvestmentByID retrieves an investment by ID
func (m *InvestmentRepository) GetInvestmentByID(id uint) (*model.Investment, error) {
	if m.MockErr != nil {
		return nil, m.MockErr
	}
	return m.MockInvestment, nil
}

// GetInvestmentsByClientID retrieves all investments for a specific client
func (m *InvestmentRepository) GetInvestmentsByClientID(clientID uint) ([]*model.Investment, error) {
	if m.MockErr != nil {
		return nil, m.MockErr
	}
	return m.MockInvestments, nil
}
