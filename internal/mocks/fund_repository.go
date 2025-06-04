package mocks

import (
	"cushon/internal/model"
)

// FundRepository is a mock implementation of repository.FundRepository
type FundRepository struct {
	CreateFundErr error
	GetAllErr     error
	MockFunds     []*model.Fund
	MockFund      *model.Fund
}

// CreateFund implements repository.FundRepository
func (m *FundRepository) CreateFund(name string) (*model.Fund, error) {
	if m.CreateFundErr != nil {
		return nil, m.CreateFundErr
	}
	return m.MockFund, nil
}

// GetAllFunds implements repository.FundRepository
func (m *FundRepository) GetAllFunds() ([]*model.Fund, error) {
	if m.GetAllErr != nil {
		return nil, m.GetAllErr
	}
	return m.MockFunds, nil
}
