package mocks

import (
	"cushon/internal/model"
)

// InvestmentService is a mock implementation of service.Investment
type InvestmentService struct {
	MockInvestment  *model.Investment
	MockInvestments []*model.Investment
	MockErr         error
}

// NewInvestment implements service.Investment
func (m *InvestmentService) NewInvestment(clientID, fundID uint, amount float32) (*model.Investment, error) {
	if m.MockErr != nil {
		return nil, m.MockErr
	}
	return m.MockInvestment, nil
}

// GetInvestment implements service.Investment
func (m *InvestmentService) GetInvestment(id uint) (*model.Investment, error) {
	if m.MockErr != nil {
		return nil, m.MockErr
	}
	return m.MockInvestment, nil
}
