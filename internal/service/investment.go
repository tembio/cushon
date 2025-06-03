package service

import (
	"cushon/internal/model"
	"cushon/internal/repository"
	"errors"
)

// Investment defines the interface for investment operations
type Investment interface {
	Create(investment *model.Investment) error
	Get(id string) (*model.Investment, error)
}

// defaultInvestmentService is a concrete implementation of InvestmentService
type defaultInvestmentService struct {
	repo repository.InvestmentRepository
}

// NewDefaultInvestmentService creates a new default investment service
func NewDefaultInvestmentService(repo repository.InvestmentRepository) *defaultInvestmentService {
	return &defaultInvestmentService{repo: repo}
}

// Create creates a new investment from a customer into a fund
func (s *defaultInvestmentService) Create(investment *model.Investment) error {
	if investment == nil {
		return errors.New("investment cannot be nil")
	}

	if investment.Amount <= 0 {
		return errors.New("investment amount must be greater than 0")
	}

	return s.repo.Create(investment)
}

// Get implements the Investment interface
func (s *defaultInvestmentService) Get(id uint) (*model.Investment, error) {
	return s.repo.GetByID(id)
}
