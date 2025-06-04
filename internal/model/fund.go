package model

// Fund represents an investment fund
type Fund struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

// FundCreate represents the data needed to create a new fund
type FundCreate struct {
	Name string `json:"name"`
}

// FundResponse represents the fund data that will be sent in API responses
type FundResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
