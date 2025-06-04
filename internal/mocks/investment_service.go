package mocks

import (
	"cushon/internal/model"
)

// InvestmentService is a mock implementation of the Investment service interface
type InvestmentService struct {
	MockInvestment  *model.Investment
	MockInvestments []*model.Investment
	MockErr         error
}

// NewInvestment creates a new investment
func (m *InvestmentService) NewInvestment(clientID, fundID uint, amount float32) (*model.Investment, error) {
	if m.MockErr != nil {
		return nil, m.MockErr
	}
	return m.MockInvestment, nil
}

// GetInvestment retrieves an investment by ID
func (m *InvestmentService) GetInvestment(id uint) (*model.Investment, error) {
	if m.MockErr != nil {
		return nil, m.MockErr
	}
	return m.MockInvestment, nil
}

// GetInvestmentsByClientID retrieves all investments for a specific client
func (m *InvestmentService) GetInvestmentsByClientID(clientID uint) ([]*model.Investment, error) {
	if m.MockErr != nil {
		return nil, m.MockErr
	}
	return m.MockInvestments, nil
}
