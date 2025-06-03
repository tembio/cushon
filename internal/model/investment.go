package model

import "time"

// Investment represents an investment in the system
type Investment struct {
	ID        int       `json:"id"`
	ClientID  int       `json:"client_id"`
	FundID    int       `json:"fund_id"`
	Amount    float64   `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// InvestmentCreate represents the data needed to create a new investment
type InvestmentCreate struct {
	ClientID int     `json:"client_id"`
	FundID   int     `json:"fund_id"`
	Amount   float64 `json:"amount"`
}

// InvestmentResponse represents the investment data that will be sent in API responses
type InvestmentResponse struct {
	ID       int     `json:"id"`
	ClientID int     `json:"client_id"`
	FundID   int     `json:"fund_id"`
	Amount   float64 `json:"amount"`
}
