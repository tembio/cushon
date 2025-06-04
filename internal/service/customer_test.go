package service

import (
	"errors"
	"testing"
	"time"

	"cushon/internal/mocks"
	"cushon/internal/model"
)

func TestDefaultCustomerService_NewRetailCustomer(t *testing.T) {
	tests := []struct {
		name         string
		customerName string
		mockCustomer *model.Customer
		mockErr      error
		wantErr      error
	}{
		{
			name:         "Valid retail customer",
			customerName: "John Doe",
			mockCustomer: &model.Customer{
				ID:         1,
				Name:       "John Doe",
				EmployerID: nil,
				CreatedAt:  time.Now(),
				UpdatedAt:  time.Now(),
			},
			mockErr: nil,
			wantErr: nil,
		},
		{
			name:         "Empty customer name",
			customerName: "",
			mockCustomer: nil,
			mockErr:      errors.New("customer name cannot be empty"),
			wantErr:      errors.New("customer name cannot be empty"),
		},
		{
			name:         "Repository error",
			customerName: "Jane Smith",
			mockCustomer: nil,
			mockErr:      errors.New("repository error"),
			wantErr:      errors.New("repository error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := &mocks.CustomerRepository{
				CreateCustomerErr: tt.mockErr,
				MockCustomer:      tt.mockCustomer,
			}

			service := NewDefaultCustomerService(mockRepo)
			got, err := service.NewRetailCustomer(tt.customerName)

			if tt.wantErr != nil {
				if err == nil || err.Error() != tt.wantErr.Error() {
					t.Errorf("NewRetailCustomer() error = %v, wantErr %v", err, tt.wantErr)
				}
				return
			}

			if err != nil {
				t.Errorf("NewRetailCustomer() unexpected error = %v", err)
				return
			}

			if got == nil {
				t.Error("NewRetailCustomer() returned nil customer")
				return
			}

			if got.Name != tt.customerName {
				t.Errorf("Name = %v, want %v", got.Name, tt.customerName)
			}

			if got.EmployerID != nil {
				t.Error("EmployerID should be nil for retail customer")
			}
		})
	}
}

func TestDefaultCustomerService_NewEmployedCustomer(t *testing.T) {
	tests := []struct {
		name         string
		customerName string
		employerID   uint
		mockCustomer *model.Customer
		mockErr      error
		wantErr      error
	}{
		{
			name:         "Valid employed customer",
			customerName: "Jane Smith",
			employerID:   1,
			mockCustomer: &model.Customer{
				ID:         1,
				Name:       "Jane Smith",
				EmployerID: uintPtr(1),
				CreatedAt:  time.Now(),
				UpdatedAt:  time.Now(),
			},
			mockErr: nil,
			wantErr: nil,
		},
		{
			name:         "Empty customer name",
			customerName: "",
			employerID:   1,
			mockCustomer: nil,
			mockErr:      errors.New("customer name cannot be empty"),
			wantErr:      errors.New("customer name cannot be empty"),
		},
		{
			name:         "Repository error",
			customerName: "John Doe",
			employerID:   2,
			mockCustomer: nil,
			mockErr:      errors.New("repository error"),
			wantErr:      errors.New("repository error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := &mocks.CustomerRepository{
				CreateCustomerErr: tt.mockErr,
				MockCustomer:      tt.mockCustomer,
			}

			service := NewDefaultCustomerService(mockRepo)
			got, err := service.NewEmployedCustomer(tt.customerName, tt.employerID)

			if tt.wantErr != nil {
				if err == nil || err.Error() != tt.wantErr.Error() {
					t.Errorf("NewEmployedCustomer() error = %v, wantErr %v", err, tt.wantErr)
				}
				return
			}

			if err != nil {
				t.Errorf("NewEmployedCustomer() unexpected error = %v", err)
				return
			}

			if got == nil {
				t.Error("NewEmployedCustomer() returned nil customer")
				return
			}

			if got.Name != tt.customerName {
				t.Errorf("Name = %v, want %v", got.Name, tt.customerName)
			}

			if got.EmployerID == nil {
				t.Error("EmployerID should not be nil for employed customer")
			} else if *got.EmployerID != tt.employerID {
				t.Errorf("EmployerID = %v, want %v", *got.EmployerID, tt.employerID)
			}
		})
	}
}

// Helper function to create a pointer to uint
func uintPtr(n uint) *uint {
	return &n
}
