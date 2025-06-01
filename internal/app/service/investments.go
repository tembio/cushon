package service

import "cushon/internal/model"

// Investment defines the interface for investment operations
type Investment interface {
	Create(investment *model.Investment) error
	Get(id string) (*model.Investment, error)
	Update(investment *model.Investment) error
	Delete(id string) error
}
