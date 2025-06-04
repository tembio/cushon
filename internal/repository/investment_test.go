package repository

import (
	"errors"
	"testing"
)

func TestInMemoryInvestmentRepository_CreateInvestment(t *testing.T) {
	tests := []struct {
		name     string
		clientID uint
		fundID   uint
		amount   float32
		wantErr  error
	}{
		{
			name:     "Valid investment",
			clientID: 1,
			fundID:   1,
			amount:   1000.0,
			wantErr:  nil,
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
					t.Errorf("CreateInvestment() error = %v, wantErr %v", err, tt.wantErr)
				}
				return
			}

			if err != nil {
				t.Errorf("CreateInvestment() unexpected error = %v", err)
				return
			}

			if got == nil {
				t.Error("CreateInvestment() returned nil investment")
				return
			}

			if got.ClientID != tt.clientID {
				t.Errorf("ClientID = %v, want %v", got.ClientID, tt.clientID)
			}
			if got.FundID != tt.fundID {
				t.Errorf("FundID = %v, want %v", got.FundID, tt.fundID)
			}
			if got.Amount != tt.amount {
				t.Errorf("Amount = %v, want %v", got.Amount, tt.amount)
			}
		})
	}
}

func TestInMemoryInvestmentRepository_GetInvestmentByID(t *testing.T) {
	tests := []struct {
		name    string
		setup   func(*InMemoryInvestmentRepository) uint
		id      uint
		wantErr error
	}{
		{
			name: "Existing investment",
			setup: func(r *InMemoryInvestmentRepository) uint {
				inv, _ := r.CreateInvestment(1, 1, 1000.0)
				return inv.ID
			},
			wantErr: nil,
		},
		{
			name: "Non-existent investment",
			setup: func(r *InMemoryInvestmentRepository) uint {
				return 999
			},
			wantErr: errors.New("investment not found"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := NewInMemoryInvestmentRepository()
			id := tt.setup(repo)

			got, err := repo.GetInvestmentByID(id)

			if tt.wantErr != nil {
				if err == nil || err.Error() != tt.wantErr.Error() {
					t.Errorf("GetInvestmentByID() error = %v, wantErr %v", err, tt.wantErr)
				}
				return
			}

			if err != nil {
				t.Errorf("GetInvestmentByID() unexpected error = %v", err)
				return
			}

			if got == nil {
				t.Error("GetInvestmentByID() returned nil investment")
				return
			}

			if got.ID != id {
				t.Errorf("ID = %v, want %v", got.ID, id)
			}
		})
	}
}

func TestInMemoryInvestmentRepository_GetInvestmentsByClientID(t *testing.T) {
	tests := []struct {
		name      string
		setup     func(*InMemoryInvestmentRepository)
		clientID  uint
		wantCount int
	}{
		{
			name: "Single investment for client",
			setup: func(r *InMemoryInvestmentRepository) {
				r.CreateInvestment(1, 1, 1000.0)
			},
			clientID:  1,
			wantCount: 1,
		},
		{
			name: "Multiple investments for client",
			setup: func(r *InMemoryInvestmentRepository) {
				r.CreateInvestment(1, 1, 1000.0)
				r.CreateInvestment(1, 2, 2000.0)
				r.CreateInvestment(1, 3, 3000.0)
			},
			clientID:  1,
			wantCount: 3,
		},
		{
			name: "No investments for client",
			setup: func(r *InMemoryInvestmentRepository) {
				r.CreateInvestment(2, 1, 1000.0)
				r.CreateInvestment(2, 2, 2000.0)
			},
			clientID:  1,
			wantCount: 0,
		},
		{
			name: "Empty repository",
			setup: func(r *InMemoryInvestmentRepository) {
				// No investments created
			},
			clientID:  1,
			wantCount: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := NewInMemoryInvestmentRepository()
			if tt.setup != nil {
				tt.setup(repo)
			}

			got, err := repo.GetInvestmentsByClientID(tt.clientID)
			if err != nil {
				t.Errorf("GetInvestmentsByClientID() error = %v", err)
				return
			}

			if len(got) != tt.wantCount {
				t.Errorf("got %d investments, want %d", len(got), tt.wantCount)
				return
			}

			// Verify all returned investments belong to the requested client
			for _, inv := range got {
				if inv.ClientID != tt.clientID {
					t.Errorf("found investment with ClientID = %v, want %v", inv.ClientID, tt.clientID)
				}
			}
		})
	}
}
