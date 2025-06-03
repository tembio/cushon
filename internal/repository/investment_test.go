package repository

import (
	"errors"
	"testing"

	"cushon/internal/model"
)

func TestInMemoryInvestmentRepository_Create(t *testing.T) {
	tests := []struct {
		name       string
		investment *model.Investment
		wantErr    error
	}{
		{
			name: "Valid investment",
			investment: &model.Investment{
				ClientID: 1,
				FundID:   1,
				Amount:   1000.0,
			},
			wantErr: nil,
		},
		{
			name:       "Nil investment",
			investment: nil,
			wantErr:    errors.New("investment cannot be nil"),
		},
		{
			name: "Invalid client ID",
			investment: &model.Investment{
				ClientID: 101,
				FundID:   1,
				Amount:   1000.0,
			},
			wantErr: errors.New("invalid user"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := NewInMemoryInvestmentRepository()
			err := repo.Create(tt.investment)

			if tt.wantErr != nil && err.Error() != tt.wantErr.Error() {
				t.Errorf("got error %v, want %v", err, tt.wantErr)
			}

			if tt.investment != nil && err == nil {
				got, err := repo.GetByID(tt.investment.ID)
				if err != nil {
					t.Errorf("failed to retrieve created investment: %v", err)
				}
				if got.ClientID != tt.investment.ClientID {
					t.Errorf("got ClientID = %v, want %v", got.ClientID, tt.investment.ClientID)
				}
				if got.FundID != tt.investment.FundID {
					t.Errorf("got FundID = %v, want %v", got.FundID, tt.investment.FundID)
				}
				if got.Amount != tt.investment.Amount {
					t.Errorf("got Amount = %v, want %v", got.Amount, tt.investment.Amount)
				}
				if got.CreatedAt.IsZero() {
					t.Error("CreatedAt was not set")
				}
				if got.UpdatedAt.IsZero() {
					t.Error("UpdatedAt was not set")
				}
			}
		})
	}
}

func TestInMemoryInvestmentRepository_GetByID(t *testing.T) {
	repo := NewInMemoryInvestmentRepository()
	testInvestment := &model.Investment{
		ClientID: 1,
		FundID:   1,
		Amount:   1000.0,
	}
	repo.Create(testInvestment)

	tests := []struct {
		name    string
		id      uint
		want    *model.Investment
		wantErr error
	}{
		{
			name:    "Existing investment",
			id:      0,
			want:    testInvestment,
			wantErr: nil,
		},
		{
			name:    "Non-existent investment",
			id:      999,
			want:    nil,
			wantErr: errors.New("investment not found"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := repo.GetByID(tt.id)

			if tt.wantErr != nil && err.Error() != tt.wantErr.Error() {
				t.Fatalf("got error %v, want %v", err, tt.wantErr)
			}

			if tt.wantErr == nil {
				if got.ClientID != tt.want.ClientID {
					t.Errorf("got ClientID = %v, want %v", got.ClientID, tt.want.ClientID)
				}
				if got.FundID != tt.want.FundID {
					t.Errorf("got FundID = %v, want %v", got.FundID, tt.want.FundID)
				}
				if got.Amount != tt.want.Amount {
					t.Errorf("got Amount = %v, want %v", got.Amount, tt.want.Amount)
				}
			}
		})
	}
}
