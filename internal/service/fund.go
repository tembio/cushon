package service

import (
	"cushon/internal/model"
	"cushon/internal/repository"
)

// Fund defines the interface for fund operations
type Fund interface {
	Create(fund *model.Fund) error
	Get(id int) (*model.Fund, error)
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
func (s *defaultFundService) Create(fund *model.Fund) error {
	return s.repo.CreateFund(fund)
}

// Get retrieves a fund by its ID
func (s *defaultFundService) Get(id int) (*model.Fund, error) {
	return s.repo.GetFundByID(id)
}
