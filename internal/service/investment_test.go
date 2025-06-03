package service

import (
	"errors"
	"testing"
	"time"

	"cushon/internal/mocks"
	"cushon/internal/model"
)

func TestDefaultInvestmentService_Create(t *testing.T) {
	tests := []struct {
		name          string
		investment    *model.Investment
		wantID        uint
		repositoryErr error
		wantErr       error
	}{
		{
			name: "Valid investment",
			investment: &model.Investment{
				ClientID: 1,
				FundID:   1,
				Amount:   1000.0,
			},
			wantID:        5,
			repositoryErr: nil,
			wantErr:       nil,
		},
		{
			name:          "Nil investment",
			investment:    nil,
			repositoryErr: nil,
			wantErr:       errors.New("investment cannot be nil"),
		},
		{
			name: "Zero amount",
			investment: &model.Investment{
				ClientID: 1,
				FundID:   1,
				Amount:   0,
			},
			repositoryErr: nil,
			wantErr:       errors.New("investment amount must be greater than 0"),
		},
		{
			name: "Negative amount",
			investment: &model.Investment{
				ClientID: 1,
				FundID:   1,
				Amount:   -100.0,
			},
			repositoryErr: nil,
			wantErr:       errors.New("investment amount must be greater than 0"),
		},
		{
			name: "Repository error",
			investment: &model.Investment{
				ClientID: 1,
				FundID:   1,
				Amount:   1000.0,
			},
			repositoryErr: errors.New("repository error"),
			wantErr:       errors.New("repository error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var mockRepo *mocks.InvestmentRepository
			now := time.Now()

			if tt.investment != nil {
				mockRepo = &mocks.InvestmentRepository{
					CreateErr: tt.repositoryErr,
					CreateInvestment: &model.Investment{
						ID:        tt.wantID,
						ClientID:  tt.investment.ClientID,
						FundID:    tt.investment.FundID,
						Amount:    tt.investment.Amount,
						CreatedAt: now,
						UpdatedAt: now,
					},
				}
			}

			service := NewDefaultInvestmentService(mockRepo)
			err := service.Create(tt.investment)

			if tt.wantErr != nil && err.Error() != tt.wantErr.Error() {
				t.Errorf("got error %v, want %v", err, tt.wantErr)
			}

			if tt.wantErr == nil && tt.investment != nil {
				if tt.investment.CreatedAt.IsZero() {
					t.Error("CreatedAt was not set")
				}
				if tt.investment.UpdatedAt.IsZero() {
					t.Error("UpdatedAt was not set")
				}
			}
		})
	}
}

func TestDefaultInvestmentService_Get(t *testing.T) {
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
				GetErr:        tt.repositoryErr,
				GetInvestment: tt.wantInvestment,
			}

			service := NewDefaultInvestmentService(mockRepo)
			gotInvestment, gotErr := service.Get(tt.ID)

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
