package repository

import (
	"cushon/internal/model"
	"errors"
)

// FundRepository defines the contract for storing and retrieving fund data.
type FundRepository interface {
	CreateFund(name string) (*model.Fund, error)
	GetAllFunds() ([]*model.Fund, error)
}

// InMemoryFundRepository is a simple in-memory implementation of FundRepository for demonstration.
type InMemoryFundRepository struct {
	funds  map[uint]*model.Fund
	nextID uint
}

// NewInMemoryFundRepository creates a new in-memory fund repository.
func NewInMemoryFundRepository() *InMemoryFundRepository {
	return &InMemoryFundRepository{
		funds:  make(map[uint]*model.Fund),
		nextID: 1,
	}
}

// GetAllFunds retrieves all funds
func (r *InMemoryFundRepository) GetAllFunds() ([]*model.Fund, error) {
	funds := make([]*model.Fund, 0, len(r.funds))
	for _, fund := range r.funds {
		funds = append(funds, fund)
	}
	return funds, nil
}

// CreateFund creates a new fund
func (r *InMemoryFundRepository) CreateFund(name string) (*model.Fund, error) {
	if name == "" {
		return nil, errors.New("fund name cannot be empty")
	}

	fund := &model.Fund{
		ID:   r.nextID,
		Name: name,
	}

	r.funds[fund.ID] = fund
	r.nextID++

	return fund, nil
}
