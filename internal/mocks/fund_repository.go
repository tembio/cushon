package mocks

import (
	"cushon/internal/model"
)

// FundRepository is a mock implementation of repository.FundRepository
type FundRepository struct {
	MockFunds []*model.Fund
	MockFund  *model.Fund
	MockErr   error
}

// CreateFund implements repository.FundRepository
func (m *FundRepository) CreateFund(name string) (*model.Fund, error) {
	if m.MockErr != nil {
		return nil, m.MockErr
	}
	return m.MockFund, nil
}

// GetAllFunds implements repository.FundRepository
func (m *FundRepository) GetAllFunds() ([]*model.Fund, error) {
	if m.MockErr != nil {
		return nil, m.MockErr
	}
	return m.MockFunds, nil
}
