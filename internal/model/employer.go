package model

// Employer represents an employer in the system
type Employer struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

// EmployerCreate represents the data needed to create a new employer
type EmployerCreate struct {
	Name string `json:"name"`
}

// EmployerResponse represents the employer data that will be sent in API responses
type EmployerResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
