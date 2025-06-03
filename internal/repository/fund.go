package repository

import "cushon/internal/model"

// FundRepository defines the contract for storing and retrieving fund data.
type FundRepository interface {
	GetFundByID(id int) (*model.Fund, error)
	CreateFund(fund *model.Fund) error
	UpdateFund(fund *model.Fund) error
	DeleteFund(id int) error
}

// InMemoryFundRepository is a simple in-memory implementation of FundRepository for demonstration.
type InMemoryFundRepository struct {
	funds map[int]*model.Fund
}

// NewInMemoryFundRepository creates a new in-memory fund repository.
func NewInMemoryFundRepository() *InMemoryFundRepository {
	return &InMemoryFundRepository{
		funds: make(map[int]*model.Fund),
	}
}

// GetFundByID retrieves a fund by its ID
func (r *InMemoryFundRepository) GetFundByID(id int) (*model.Fund, error) {
	if fund, exists := r.funds[id]; exists {
		return fund, nil
	}
	return nil, nil
}

// CreateFund creates a new fund
func (r *InMemoryFundRepository) CreateFund(fund *model.Fund) error {
	r.funds[fund.ID] = fund
	return nil
}

// UpdateFund updates an existing fund
func (r *InMemoryFundRepository) UpdateFund(fund *model.Fund) error {
	r.funds[fund.ID] = fund
	return nil
}

// DeleteFund deletes a fund by its ID
func (r *InMemoryFundRepository) DeleteFund(id int) error {
	delete(r.funds, id)
	return nil
}
