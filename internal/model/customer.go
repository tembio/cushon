package model

import "time"

// Customer represents a user in the system
type Customer struct {
	ID         uint      `json:"id"`
	Name       string    `json:"name"`
	EmployerID *uint     `json:"employer_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// CustomerCreate represents the data needed to create a new customer
type CustomerCreate struct {
	Name       string `json:"name"`
	EmployerID *uint  `json:"employer_id"`
}

// CustomerResponse represents the customer data that will be sent in API responses
type CustomerResponse struct {
	ID         uint   `json:"id"`
	Name       string `json:"name"`
	EmployerID *uint  `json:"employer_id,omitempty"`
}
