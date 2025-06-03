package repository

import "cushon/internal/model"

// InvestmentRepository defines the contract for storing and retrieving investment data.
// This interface allows us to swap out different storage implementations (e.g., in-memory, SQL...).
type InvestmentRepository interface {
	GetInvestmentByID(id int) (*model.Investment, error)
}

// inMemoryInvestmentRepository is a simple in-memory implementation of InvestmentRepository for demonstration.
type inMemoryInvestmentRepository struct {
	investments map[int]*model.Investment
}

// NewInMemoryInvestmentRepository creates a new in-memory investment repository.
func NewInMemoryInvestmentRepository() *inMemoryInvestmentRepository {
	return &inMemoryInvestmentRepository{
		investments: make(map[int]*model.Investment),
	}
}
