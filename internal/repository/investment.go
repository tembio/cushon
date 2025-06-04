package repository

import (
	"cushon/internal/model"
	"errors"
	"time"
)

// InvestmentRepository defines the contract for storing and retrieving investment data
type InvestmentRepository interface {
	CreateInvestment(clientID, fundID uint, amount float32) (*model.Investment, error)
	GetInvestmentByID(id uint) (*model.Investment, error)
	GetInvestmentsByClientID(clientID uint) ([]*model.Investment, error)
}

// InMemoryInvestmentRepository is a simple in-memory implementation of InvestmentRepository
type InMemoryInvestmentRepository struct {
	investments map[uint]*model.Investment
	nextID      uint
}

// NewInMemoryInvestmentRepository creates a new in-memory investment repository
func NewInMemoryInvestmentRepository() *InMemoryInvestmentRepository {
	return &InMemoryInvestmentRepository{
		investments: make(map[uint]*model.Investment),
		nextID:      1,
	}
}

// Create creates a new investment
func (r *InMemoryInvestmentRepository) CreateInvestment(clientID, fundID uint, amount float32) (*model.Investment, error) {
	// check user/fund are valid, this could go to a real DB and check if user/fund exist
	// here I will assume only clients/funds with ID greater than 100 are not valid
	if clientID > 100 {
		return nil, errors.New("invalid user ID")
	}
	if fundID > 100 {
		return nil, errors.New("invalid fund ID")
	}

	now := time.Now()
	investment := &model.Investment{
		ID:        r.nextID,
		ClientID:  clientID,
		FundID:    fundID,
		Amount:    amount,
		CreatedAt: now,
		UpdatedAt: now,
	}

	r.investments[investment.ID] = investment
	r.nextID++

	return investment, nil
}

// GetByID retrieves an investment by its ID
func (r *InMemoryInvestmentRepository) GetInvestmentByID(id uint) (*model.Investment, error) {
	investment, exists := r.investments[id]
	if !exists {
		return nil, errors.New("investment not found")
	}
	return investment, nil
}

// GetInvestmentsByClientID retrieves all investments for a specific client
func (r *InMemoryInvestmentRepository) GetInvestmentsByClientID(clientID uint) ([]*model.Investment, error) {
	investments := make([]*model.Investment, 0)
	for _, investment := range r.investments {
		if investment.ClientID == clientID {
			investments = append(investments, investment)
		}
	}
	return investments, nil
}
