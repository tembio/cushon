package service

import (
	"errors"
	"testing"
	"time"

	"cushon/internal/mocks"
	"cushon/internal/model"
)

func TestDefaultInvestmentService_NewInvestment(t *testing.T) {
	tests := []struct {
		name             string
		clientID         uint
		fundID           uint
		amount           float32
		wantInvestmentID uint
		repositoryErr    error
		wantErr          error
	}{
		{
			name:             "Valid investment",
			clientID:         1,
			fundID:           1,
			amount:           1000.0,
			wantInvestmentID: 5,
			wantErr:          nil,
		},
		{
			name:     "Zero amount",
			clientID: 1,
			fundID:   1,
			amount:   0,
			wantErr:  errors.New("investment amount must be greater than 0"),
		},
		{
			name:     "Negative amount",
			clientID: 1,
			fundID:   1,
			amount:   -100.0,
			wantErr:  errors.New("investment amount must be greater than 0"),
		},
		{
			name:          "Repository error",
			clientID:      1,
			fundID:        1,
			amount:        1000.0,
			repositoryErr: errors.New("repository error"),
			wantErr:       errors.New("repository error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var mockRepo *mocks.InvestmentRepository
			now := time.Now()

			if tt.wantErr == nil || tt.repositoryErr != nil {
				mockRepo = &mocks.InvestmentRepository{
					MockErr: tt.repositoryErr,
					MockInvestment: &model.Investment{
						ID:        tt.wantInvestmentID,
						ClientID:  tt.clientID,
						FundID:    tt.fundID,
						Amount:    tt.amount,
						CreatedAt: now,
						UpdatedAt: now,
					},
				}
			}

			service := NewDefaultInvestmentService(mockRepo)
			gotInvestment, err := service.NewInvestment(tt.clientID, tt.fundID, tt.amount)

			if tt.wantErr != nil {
				if err == nil || err.Error() != tt.wantErr.Error() {
					t.Errorf("got error %v, want %v", err, tt.wantErr)
				}
				return
			}

			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}

			if gotInvestment == nil {
				t.Error("expected investment to be returned")
				return
			}

			if gotInvestment.ClientID != tt.clientID {
				t.Errorf("got ClientID %v, want %v", gotInvestment.ClientID, tt.clientID)
			}
			if gotInvestment.FundID != tt.fundID {
				t.Errorf("got FundID %v, want %v", gotInvestment.FundID, tt.fundID)
			}
			if gotInvestment.Amount != tt.amount {
				t.Errorf("got Amount %v, want %v", gotInvestment.Amount, tt.amount)
			}
		})
	}
}

func TestDefaultInvestmentService_GetInvestment(t *testing.T) {
	tests := []struct {
		name           string
		ID             uint
		wantInvestment *model.Investment
		repositoryErr  error
	}{
		{
			name: "Valid investment",
			ID:   5,
			wantInvestment: &model.Investment{
				ID:       5,
				ClientID: 1,
				FundID:   1,
				Amount:   1000.0,
			},
			repositoryErr: nil,
		},
		{
			name:           "Repository error",
			ID:             1,
			wantInvestment: nil,
			repositoryErr:  errors.New("repository error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := &mocks.InvestmentRepository{
				MockErr:        tt.repositoryErr,
				MockInvestment: tt.wantInvestment,
			}

			service := NewDefaultInvestmentService(mockRepo)
			gotInvestment, gotErr := service.GetInvestment(tt.ID)

			if tt.repositoryErr != nil && gotErr.Error() != tt.repositoryErr.Error() {
				t.Errorf("got error %v, want %v", gotErr, tt.repositoryErr)
			}

			if tt.repositoryErr == nil && tt.wantInvestment != nil {
				if tt.wantInvestment.ID != gotInvestment.ID {
					t.Errorf("Incorrect investment retrieved, got %v, want %v", gotInvestment, tt.wantInvestment)
				}
			}
		})
	}
}

func TestDefaultInvestmentService_GetInvestmentsByClientID(t *testing.T) {
	tests := []struct {
		name            string
		clientID        uint
		wantInvestments []*model.Investment
		repositoryErr   error
	}{
		{
			name:     "Single investment",
			clientID: 1,
			wantInvestments: []*model.Investment{
				{
					ID:       1,
					ClientID: 1,
					FundID:   1,
					Amount:   1000.0,
				},
			},
			repositoryErr: nil,
		},
		{
			name:     "Multiple investments",
			clientID: 1,
			wantInvestments: []*model.Investment{
				{
					ID:       1,
					ClientID: 1,
					FundID:   1,
					Amount:   1000.0,
				},
				{
					ID:       2,
					ClientID: 1,
					FundID:   2,
					Amount:   2000.0,
				},
			},
			repositoryErr: nil,
		},
		{
			name:            "No investments",
			clientID:        1,
			wantInvestments: []*model.Investment{},
			repositoryErr:   nil,
		},
		{
			name:            "Repository error",
			clientID:        1,
			wantInvestments: nil,
			repositoryErr:   errors.New("repository error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := &mocks.InvestmentRepository{
				MockErr:         tt.repositoryErr,
				MockInvestments: tt.wantInvestments,
			}

			service := NewDefaultInvestmentService(mockRepo)
			gotInvestments, gotErr := service.GetInvestmentsByClientID(tt.clientID)

			if tt.repositoryErr != nil && gotErr.Error() != tt.repositoryErr.Error() {
				t.Errorf("got error %v, want %v", gotErr, tt.repositoryErr)
			}

			if tt.repositoryErr == nil {
				if len(gotInvestments) != len(tt.wantInvestments) {
					t.Errorf("got %d investments, want %d", len(gotInvestments), len(tt.wantInvestments))
					return
				}

				for i, want := range tt.wantInvestments {
					if gotInvestments[i].ID != want.ID {
						t.Errorf("investment[%d].ID = %v, want %v", i, gotInvestments[i].ID, want.ID)
					}
					if gotInvestments[i].ClientID != want.ClientID {
						t.Errorf("investment[%d].ClientID = %v, want %v", i, gotInvestments[i].ClientID, want.ClientID)
					}
					if gotInvestments[i].FundID != want.FundID {
						t.Errorf("investment[%d].FundID = %v, want %v", i, gotInvestments[i].FundID, want.FundID)
					}
					if gotInvestments[i].Amount != want.Amount {
						t.Errorf("investment[%d].Amount = %v, want %v", i, gotInvestments[i].Amount, want.Amount)
					}
				}
			}
		})
	}
}
