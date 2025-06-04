package mocks

import (
	"cushon/internal/model"
)

// FundService is a mock implementation of service.Fund
type FundService struct {
	MockFund  *model.Fund
	MockFunds []*model.Fund
	MockErr   error
}

// NewFund implements service.Fund
func (m *FundService) NewFund(name string) (*model.Fund, error) {
	if m.MockErr != nil {
		return nil, m.MockErr
	}
	return m.MockFund, nil
}

// GetAllFunds implements service.Fund
func (m *FundService) GetAllFunds() ([]*model.Fund, error) {
	if m.MockErr != nil {
		return nil, m.MockErr
	}
	return m.MockFunds, nil
}
