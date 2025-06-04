package repository

import (
	"errors"
	"testing"

	"cushon/internal/model"
)

func TestInMemoryEmployerRepository_Create(t *testing.T) {
	tests := []struct {
		name    string
		empName string
		wantEmp *model.Employer
		wantErr error
	}{
		{
			name:    "Valid employer",
			empName: "Test Employer",
			wantEmp: &model.Employer{
				ID:   1,
				Name: "Test Employer",
			},
			wantErr: nil,
		},
		{
			name:    "Empty employer name",
			empName: "",
			wantEmp: nil,
			wantErr: errors.New("employer name cannot be empty"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := NewInMemoryEmployerRepository()
			got, err := repo.CreateEmployer(tt.empName)

			if tt.wantErr != nil {
				if err == nil || err.Error() != tt.wantErr.Error() {
					t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				}
				return
			}

			if err != nil {
				t.Errorf("Create() unexpected error = %v", err)
				return
			}

			if got == nil {
				t.Error("Create() returned nil employer")
				return
			}

			if got.ID != tt.wantEmp.ID {
				t.Errorf("ID = %v, want %v", got.ID, tt.wantEmp.ID)
			}
			if got.Name != tt.wantEmp.Name {
				t.Errorf("Name = %v, want %v", got.Name, tt.wantEmp.Name)
			}

			// Verify the employer was stored
			stored, exists := repo.employers[got.ID]
			if !exists {
				t.Error("employer was not stored")
				return
			}
			if stored.ID != got.ID || stored.Name != got.Name {
				t.Error("stored employer does not match created employer")
			}
		})
	}
}
