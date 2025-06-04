package service

import (
	"cushon/internal/model"
	"cushon/internal/repository"
	"errors"
)

// Investment defines the interface for investment operations
type Investment interface {
	NewInvestment(clientID, fundID uint, amount float32) (*model.Investment, error)
	GetInvestment(id uint) (*model.Investment, error)
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
func (s *defaultInvestmentService) NewInvestment(clientID, fundID uint, amount float32) (*model.Investment, error) {
	if amount <= 0 {
		return nil, errors.New("investment amount must be greater than 0")
	}

	return s.repo.CreateInvestment(clientID, fundID, amount)
}

// GetInvestment implements the Investment interface
func (s *defaultInvestmentService) GetInvestment(id uint) (*model.Investment, error) {
	return s.repo.GetInvestmentByID(id)
}
