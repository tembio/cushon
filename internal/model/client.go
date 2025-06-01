package model

import "time"

// Client represents a user in the system
type Client struct {
	ID         int       `json:"id"`
	Email      string    `json:"email"`
	Name       string    `json:"name"`
	EmployerID int       `json:"employer_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// ClientCreate represents the data needed to create a new client
type ClientCreate struct {
	Email      string `json:"email" validate:"required,email"`
	Name       string `json:"name" validate:"required"`
	EmployerID int    `json:"employer_id" validate:"required"`
}

// ClientResponse represents the client data that will be sent in API responses
type ClientResponse struct {
	ID         int    `json:"id"`
	Email      string `json:"email"`
	Name       string `json:"name"`
	EmployerID int    `json:"employer_id"`
}
