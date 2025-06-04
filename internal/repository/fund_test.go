package repository

import (
	"errors"
	"testing"

	"cushon/internal/model"
)

func TestInMemoryFundRepository_GetAllFunds(t *testing.T) {
	tests := []struct {
		name      string
		setup     func(*InMemoryFundRepository)
		wantFunds []*model.Fund
	}{
		{
			name:      "Empty repository",
			setup:     func(r *InMemoryFundRepository) {},
			wantFunds: []*model.Fund{},
		},
		{
			name: "Single fund",
			setup: func(r *InMemoryFundRepository) {
				r.CreateFund("Test Fund 1")
			},
			wantFunds: []*model.Fund{
				{
					ID:   1,
					Name: "Test Fund 1",
				},
			},
		},
		{
			name: "Multiple funds",
			setup: func(r *InMemoryFundRepository) {
				r.CreateFund("Test Fund 1")
				r.CreateFund("Test Fund 2")
			},
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
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := NewInMemoryFundRepository()
			if tt.setup != nil {
				tt.setup(repo)
			}

			got, err := repo.GetAllFunds()
			if err != nil {
				t.Errorf("GetAllFunds() error = %v", err)
				return
			}

			if len(got) != len(tt.wantFunds) {
				t.Errorf("got %d funds, want %d", len(got), len(tt.wantFunds))
				return
			}

			for i, wantFund := range tt.wantFunds {
				if got[i].ID != wantFund.ID {
					t.Errorf("fund[%d].ID = %v, want %v", i, got[i].ID, wantFund.ID)
				}
				if got[i].Name != wantFund.Name {
					t.Errorf("fund[%d].Name = %v, want %v", i, got[i].Name, wantFund.Name)
				}
			}
		})
	}
}

func TestInMemoryFundRepository_CreateFund(t *testing.T) {
	tests := []struct {
		name     string
		fundName string
		wantErr  error
	}{
		{
			name:     "Valid fund",
			fundName: "Test Fund",
			wantErr:  nil,
		},
		{
			name:     "Empty fund name",
			fundName: "",
			wantErr:  errors.New("fund name cannot be empty"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := NewInMemoryFundRepository()
			got, err := repo.CreateFund(tt.fundName)

			if tt.wantErr != nil {
				if err == nil || err.Error() != tt.wantErr.Error() {
					t.Errorf("CreateFund() error = %v, wantErr %v", err, tt.wantErr)
				}
				return
			}

			if err != nil {
				t.Errorf("CreateFund() unexpected error = %v", err)
				return
			}

			if got == nil {
				t.Error("CreateFund() returned nil fund")
				return
			}

			// Verify the fund was stored
			funds, err := repo.GetAllFunds()
			if err != nil {
				t.Errorf("GetAllFunds() error = %v", err)
				return
			}

			found := false
			for _, fund := range funds {
				if fund.ID == got.ID && fund.Name == tt.fundName {
					found = true
					break
				}
			}
			if !found {
				t.Error("fund was not stored correctly")
			}
		})
	}
}
