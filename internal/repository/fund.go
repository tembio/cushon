package repository

import "cushon/internal/model"

// FundRepository defines the contract for storing and retrieving fund data.
// This interface allows us to swap out different storage implementations (e.g., in-memory, SQL...).
type FundRepository interface {
	GetFundByID(id int) (*model.Fund, error)
	CreateFund(fund *model.Fund) error
	UpdateFund(fund *model.Fund) error
	DeleteFund(id int) error
}

// inMemoryFundRepository is a simple in-memory implementation of FundRepository for demonstration.
type inMemoryFundRepository struct {
	funds map[int]*model.Fund
}

// NewInMemoryFundRepository creates a new in-memory fund repository.
func NewInMemoryFundRepository() *inMemoryFundRepository {
	return &inMemoryFundRepository{
		funds: make(map[int]*model.Fund),
	}
}

// GetFundByID retrieves a fund by its ID
func (r *inMemoryFundRepository) GetFundByID(id int) (*model.Fund, error) {
	if fund, exists := r.funds[id]; exists {
		return fund, nil
	}
	return nil, nil
}

// CreateFund creates a new fund
func (r *inMemoryFundRepository) CreateFund(fund *model.Fund) error {
	r.funds[fund.ID] = fund
	return nil
}

// UpdateFund updates an existing fund
func (r *inMemoryFundRepository) UpdateFund(fund *model.Fund) error {
	r.funds[fund.ID] = fund
	return nil
}

// DeleteFund deletes a fund by its ID
func (r *inMemoryFundRepository) DeleteFund(id int) error {
	delete(r.funds, id)
	return nil
}
