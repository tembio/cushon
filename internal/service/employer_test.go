package service

import (
	"errors"
	"testing"

	"cushon/internal/mocks"
	"cushon/internal/model"
)

func TestDefaultEmployerService_NewEmployer(t *testing.T) {
	tests := []struct {
		name          string
		employerName  string
		mockEmployer  *model.Employer
		repositoryErr error
		wantErr       error
	}{
		{
			name:         "Valid employer",
			employerName: "Test Company",
			mockEmployer: &model.Employer{
				ID:   1,
				Name: "Test Company",
			},
			repositoryErr: nil,
			wantErr:       nil,
		},
		{
			name:          "Empty employer name",
			employerName:  "",
			mockEmployer:  nil,
			repositoryErr: errors.New("employer name cannot be empty"),
			wantErr:       errors.New("employer name cannot be empty"),
		},
		{
			name:          "Repository error",
			employerName:  "Test Company",
			mockEmployer:  nil,
			repositoryErr: errors.New("repository error"),
			wantErr:       errors.New("repository error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := &mocks.EmployerRepository{
				MockErr:      tt.repositoryErr,
				MockEmployer: tt.mockEmployer,
			}

			service := NewDefaultEmployerService(mockRepo)
			got, err := service.NewEmployer(tt.employerName)

			if tt.wantErr != nil {
				if err == nil || err.Error() != tt.wantErr.Error() {
					t.Errorf("NewEmployer() error = %v, wantErr %v", err, tt.wantErr)
				}
				return
			}

			if err != nil {
				t.Errorf("NewEmployer() unexpected error = %v", err)
				return
			}

			if got == nil {
				t.Error("NewEmployer() returned nil employer")
				return
			}

			if got.Name != tt.employerName {
				t.Errorf("Name = %v, want %v", got.Name, tt.employerName)
			}

			if got.ID != tt.mockEmployer.ID {
				t.Errorf("ID = %v, want %v", got.ID, tt.mockEmployer.ID)
			}
		})
	}
}
