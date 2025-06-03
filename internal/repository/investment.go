package repository

import (
	"cushon/internal/model"
	"errors"
	"time"
)

// InvestmentRepository defines the interface for investment storage operations
type InvestmentRepository interface {
	Create(investment *model.Investment) error
	GetByID(id uint) (*model.Investment, error)
}

// InMemoryInvestmentRepository is a simple in-memory implementation of InvestmentRepository
type InMemoryInvestmentRepository struct {
	investmentIDCounter uint
	investments         map[uint]*model.Investment
}

// NewInMemoryInvestmentRepository creates a new in-memory investment repository
func NewInMemoryInvestmentRepository() *InMemoryInvestmentRepository {
	return &InMemoryInvestmentRepository{
		investments: make(map[uint]*model.Investment),
	}
}

// Create adds a new investment to the repository
func (r *InMemoryInvestmentRepository) Create(investment *model.Investment) error {
	if investment == nil {
		return errors.New("investment cannot be nil")
	}

	// check users are valid, this could go to a real DB and check if user exists
	// here I will assume only clients with ID below 100 are valid
	if investment.ClientID > 100 {
		return errors.New("invalid user")
	}

	investment.ID = r.investmentIDCounter
	investment.CreatedAt = time.Now()
	investment.UpdatedAt = time.Now()
	r.investments[investment.ID] = investment

	r.investmentIDCounter++

	return nil
}

// GetByID retrieves an investment from the repository
func (r *InMemoryInvestmentRepository) GetByID(id uint) (*model.Investment, error) {
	investment, exists := r.investments[id]
	if !exists {
		return nil, errors.New("investment not found")
	}
	return investment, nil
}
