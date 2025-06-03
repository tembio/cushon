package service

import (
	"cushon/internal/model"
)

// Employer defines the interface for employer operations
type Employer interface {
	Create(employer *model.Employer) error
}
