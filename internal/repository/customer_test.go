package repository

import (
	"errors"
	"testing"
)

func TestInMemoryCustomerRepository_CreateCustomer(t *testing.T) {
	tests := []struct {
		name         string
		customerName string
		employerID   *uint
		wantErr      error
	}{
		{
			name:         "Valid retail customer",
			customerName: "John Doe",
			employerID:   nil,
			wantErr:      nil,
		},
		{
			name:         "Valid employed customer",
			customerName: "Jane Smith",
			employerID:   uintPtr(1),
			wantErr:      nil,
		},
		{
			name:         "Empty customer name",
			customerName: "",
			employerID:   nil,
			wantErr:      errors.New("customer name cannot be empty"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := NewInMemoryCustomerRepository()
			got, err := repo.CreateCustomer(tt.customerName, tt.employerID)

			if tt.wantErr != nil {
				if err == nil || err.Error() != tt.wantErr.Error() {
					t.Errorf("CreateCustomer() error = %v, wantErr %v", err, tt.wantErr)
				}
				return
			}

			if err != nil {
				t.Errorf("CreateCustomer() unexpected error = %v", err)
				return
			}

			if got == nil {
				t.Error("CreateCustomer() returned nil customer")
				return
			}

			if got.Name != tt.customerName {
				t.Errorf("Name = %v, want %v", got.Name, tt.customerName)
			}

			if tt.employerID == nil {
				if got.EmployerID != nil {
					t.Error("EmployerID should be nil for retail customer")
				}
			} else {
				if got.EmployerID == nil {
					t.Error("EmployerID should not be nil for employed customer")
				} else if uint(*got.EmployerID) != *tt.employerID {
					t.Errorf("EmployerID = %v, want %v", *got.EmployerID, *tt.employerID)
				}
			}

			stored, exists := repo.customers[got.ID]
			if !exists {
				t.Error("customer was not stored in the repository")
				return
			}

			if stored.ID != got.ID {
				t.Errorf("stored customer ID = %v, want %v", stored.ID, got.ID)
			}
			if stored.Name != got.Name {
				t.Errorf("stored customer Name = %v, want %v", stored.Name, got.Name)
			}
			if (stored.EmployerID == nil) != (got.EmployerID == nil) {
				t.Error("stored customer EmployerID nil status does not match")
			} else if stored.EmployerID != nil && got.EmployerID != nil && *stored.EmployerID != *got.EmployerID {
				t.Errorf("stored customer EmployerID = %v, want %v", *stored.EmployerID, *got.EmployerID)
			}
		})
	}
}

// Helper function to create a pointer to uint
func uintPtr(n uint) *uint {
	return &n
}
