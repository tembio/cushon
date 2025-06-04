package service

import (
	"cushon/internal/model"
	"cushon/internal/repository"
)

// Fund defines the interface for fund operations
type Fund interface {
	NewFund(name string) (*model.Fund, error)
	GetAllFunds() ([]*model.Fund, error)
}

// defaultFundService is a concrete implementation of FundService
type defaultFundService struct {
	repo repository.FundRepository
}

// NewDefaultFundService creates a new default fund service
func NewDefaultFundService(repo repository.FundRepository) *defaultFundService {
	return &defaultFundService{repo: repo}
}

// Create creates a new fund
func (s *defaultFundService) NewFund(name string) (*model.Fund, error) {
	return s.repo.CreateFund(name)
}

// GetAllFunds retrieves all funds
func (s *defaultFundService) GetAllFunds() ([]*model.Fund, error) {
	return s.repo.GetAllFunds()
}
