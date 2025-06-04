package service

import (
	"errors"
	"testing"

	"cushon/internal/mocks"
	"cushon/internal/model"
)

func TestDefaultFundService_NewFund(t *testing.T) {
	tests := []struct {
		name          string
		fundName      string
		wantFund      *model.Fund
		repositoryErr error
		wantErr       error
	}{
		{
			name:     "Valid fund",
			fundName: "Test Fund",
			wantFund: &model.Fund{
				ID:   1,
				Name: "Test Fund",
			},
			wantErr: nil,
		},
		{
			name:          "Empty fund name",
			fundName:      "",
			wantFund:      nil,
			repositoryErr: errors.New("fund name cannot be empty"),
			wantErr:       errors.New("fund name cannot be empty"),
		},
		{
			name:          "Repository error",
			fundName:      "Test Fund",
			wantFund:      nil,
			repositoryErr: errors.New("repository error"),
			wantErr:       errors.New("repository error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := &mocks.FundRepository{
				CreateFundErr: tt.repositoryErr,
				MockFund:      tt.wantFund,
			}

			service := NewDefaultFundService(mockRepo)
			got, err := service.NewFund(tt.fundName)

			if tt.wantErr != nil {
				if err == nil || err.Error() != tt.wantErr.Error() {
					t.Errorf("NewFund() error = %v, wantErr %v", err, tt.wantErr)
				}
				return
			}

			if err != nil {
				t.Errorf("NewFund() unexpected error = %v", err)
				return
			}

			if got == nil {
				t.Error("NewFund() returned nil fund")
				return
			}

			if got.ID != tt.wantFund.ID {
				t.Errorf("ID = %v, want %v", got.ID, tt.wantFund.ID)
			}
			if got.Name != tt.wantFund.Name {
				t.Errorf("Name = %v, want %v", got.Name, tt.wantFund.Name)
			}
		})
	}
}

func TestDefaultFundService_GetAllFunds(t *testing.T) {
	tests := []struct {
		name          string
		wantFunds     []*model.Fund
		repositoryErr error
		wantErr       error
	}{
		{
			name: "Multiple funds",
			wantFunds: []*model.Fund{
				{
					ID:   1,
					Name: "Test Fund 1",
				},
				{
					ID:   2,
					Name: "Test Fund 2",
				},
			},
			wantErr: nil,
		},
		{
			name:      "Empty funds",
			wantFunds: []*model.Fund{},
			wantErr:   nil,
		},
		{
			name:          "Repository error",
			wantFunds:     nil,
			repositoryErr: errors.New("repository error"),
			wantErr:       errors.New("repository error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := &mocks.FundRepository{
				GetAllErr: tt.repositoryErr,
				MockFunds: tt.wantFunds,
			}

			service := NewDefaultFundService(mockRepo)
			gotFunds, err := service.GetAllFunds()

			if tt.wantErr != nil {
				if err == nil || err.Error() != tt.wantErr.Error() {
					t.Errorf("GetAllFunds() error = %v, wantErr %v", err, tt.wantErr)
				}
				return
			}

			if err != nil {
				t.Errorf("GetAllFunds() unexpected error = %v", err)
				return
			}

			if len(gotFunds) != len(tt.wantFunds) {
				t.Errorf("got %d funds, want %d", len(gotFunds), len(tt.wantFunds))
				return
			}

			for i, wantFund := range tt.wantFunds {
				if gotFunds[i].ID != wantFund.ID {
					t.Errorf("fund[%d].ID = %v, want %v", i, gotFunds[i].ID, wantFund.ID)
				}
				if gotFunds[i].Name != wantFund.Name {
					t.Errorf("fund[%d].Name = %v, want %v", i, gotFunds[i].Name, wantFund.Name)
				}
			}
		})
	}
}
