package repository

import (
	"errors"
	"testing"

	"cushon/internal/model"
)

func TestInMemoryInvestmentRepository_Create(t *testing.T) {
	tests := []struct {
		name      string
		clientID  uint
		fundID    uint
		amount    float32
		wantErr   error
		checkTime bool
	}{
		{
			name:      "Valid investment",
			clientID:  1,
			fundID:    1,
			amount:    1000.0,
			wantErr:   nil,
			checkTime: true,
		},
		{
			name:     "Invalid client ID",
			clientID: 101,
			fundID:   1,
			amount:   1000.0,
			wantErr:  errors.New("invalid user ID"),
		},
		{
			name:     "Invalid fund ID",
			clientID: 1,
			fundID:   101,
			amount:   1000.0,
			wantErr:  errors.New("invalid fund ID"),
		},
		{
			name:     "Zero amount",
			clientID: 1,
			fundID:   1,
			amount:   0,
			wantErr:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := NewInMemoryInvestmentRepository()
			got, err := repo.CreateInvestment(tt.clientID, tt.fundID, tt.amount)

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
				t.Error("Create() returned nil investment")
				return
			}

			// Verify the investment was stored
			stored, err := repo.GetInvestmentByID(got.ID)
			if err != nil {
				t.Errorf("GetByID() error = %v", err)
				return
			}

			if stored.ClientID != tt.clientID {
				t.Errorf("ClientID = %v, want %v", stored.ClientID, tt.clientID)
			}
			if stored.FundID != tt.fundID {
				t.Errorf("FundID = %v, want %v", stored.FundID, tt.fundID)
			}
			if stored.Amount != tt.amount {
				t.Errorf("Amount = %v, want %v", stored.Amount, tt.amount)
			}

			if tt.checkTime {
				if stored.CreatedAt.IsZero() {
					t.Error("CreatedAt was not set")
				}
				if stored.UpdatedAt.IsZero() {
					t.Error("UpdatedAt was not set")
				}
				if stored.CreatedAt != stored.UpdatedAt {
					t.Error("CreatedAt and UpdatedAt should be equal for new investments")
				}
			}
		})
	}
}

func TestInMemoryInvestmentRepository_GetByID(t *testing.T) {
	tests := []struct {
		name    string
		setup   func(*InMemoryInvestmentRepository)
		id      uint
		want    *model.Investment
		wantErr error
	}{
		{
			name: "Existing investment",
			setup: func(r *InMemoryInvestmentRepository) {
				r.CreateInvestment(1, 1, 1000.0)
			},
			id: 1,
			want: &model.Investment{
				ID:       1,
				ClientID: 1,
				FundID:   1,
				Amount:   1000.0,
			},
			wantErr: nil,
		},
		{
			name: "Non-existent investment",
			setup: func(r *InMemoryInvestmentRepository) {
				// No setup needed
			},
			id:      999,
			want:    nil,
			wantErr: errors.New("investment not found"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := NewInMemoryInvestmentRepository()
			if tt.setup != nil {
				tt.setup(repo)
			}

			got, err := repo.GetInvestmentByID(tt.id)

			if tt.wantErr != nil {
				if err == nil || err.Error() != tt.wantErr.Error() {
					t.Errorf("GetByID() error = %v, wantErr %v", err, tt.wantErr)
				}
				return
			}

			if err != nil {
				t.Errorf("GetByID() unexpected error = %v", err)
				return
			}

			if got.ClientID != tt.want.ClientID {
				t.Errorf("ClientID = %v, want %v", got.ClientID, tt.want.ClientID)
			}
			if got.FundID != tt.want.FundID {
				t.Errorf("FundID = %v, want %v", got.FundID, tt.want.FundID)
			}
			if got.Amount != tt.want.Amount {
				t.Errorf("Amount = %v, want %v", got.Amount, tt.want.Amount)
			}
		})
	}
}
