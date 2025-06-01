package service

import (
	"cushon/internal/app/repository"
	"cushon/internal/model"
)

// Client defines the interface for client operations
type Client interface {
	Create(client *model.Client) error
	Get(id int) (*model.Client, error)
	Update(client *model.Client) error
	Delete(id int) error
}

// defaultClientService is a concrete implementation of ClientService.
type defaultClientService struct {
	repo repository.ClientRepository
}

// NewDefaultClientService creates a new default user service.
func NewDefaultClientService(repo repository.ClientRepository) *defaultClientService {
	return &defaultClientService{repo: repo}
}
