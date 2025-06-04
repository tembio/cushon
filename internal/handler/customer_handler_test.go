package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"cushon/internal/mocks"
	"cushon/internal/model"
)

func TestCustomerHandler_Create(t *testing.T) {
	now := time.Now()
	tests := []struct {
		name           string
		requestBody    model.CustomerCreate
		mockCustomer   *model.Customer
		mockErr        error
		expectedStatus int
		expectedBody   model.CustomerResponse
	}{
		{
			name: "Create retail customer successfully",
			requestBody: model.CustomerCreate{
				Name: "John Doe",
			},
			mockCustomer: &model.Customer{
				ID:         1,
				Name:       "John Doe",
				EmployerID: nil,
				CreatedAt:  now,
				UpdatedAt:  now,
			},
			mockErr:        nil,
			expectedStatus: http.StatusCreated,
			expectedBody: model.CustomerResponse{
				ID:         1,
				Name:       "John Doe",
				EmployerID: nil,
			},
		},
		{
			name: "Create employed customer successfully",
			requestBody: model.CustomerCreate{
				Name:       "Jane Smith",
				EmployerID: uintPtr(1),
			},
			mockCustomer: &model.Customer{
				ID:         2,
				Name:       "Jane Smith",
				EmployerID: uintPtr(1),
				CreatedAt:  now,
				UpdatedAt:  now,
			},
			mockErr:        nil,
			expectedStatus: http.StatusCreated,
			expectedBody: model.CustomerResponse{
				ID:         2,
				Name:       "Jane Smith",
				EmployerID: uintPtr(1),
			},
		},
		{
			name: "Empty customer name",
			requestBody: model.CustomerCreate{
				Name: "",
			},
			mockCustomer:   nil,
			mockErr:        errors.New("customer name cannot be empty"),
			expectedStatus: http.StatusBadRequest,
			expectedBody:   model.CustomerResponse{},
		},
		{
			name: "Invalid request body",
			requestBody: model.CustomerCreate{
				Name: "Invalid JSON",
			},
			mockCustomer:   nil,
			mockErr:        nil,
			expectedStatus: http.StatusBadRequest,
			expectedBody:   model.CustomerResponse{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockService := &mocks.CustomerService{
				MockCustomer: tt.mockCustomer,
				MockErr:      tt.mockErr,
			}

			handler := NewCustomerHandler(mockService)

			var req *http.Request
			if tt.name == "Invalid request body" {
				// Send invalid JSON
				req = httptest.NewRequest("POST", "/customers", bytes.NewBufferString("invalid json"))
			} else {
				body, _ := json.Marshal(tt.requestBody)
				req = httptest.NewRequest("POST", "/customers", bytes.NewBuffer(body))
			}

			rr := httptest.NewRecorder()

			handler.Create(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v",
					rr.Code, tt.expectedStatus)
			}

			// For successful requests, check response body
			if tt.expectedStatus == http.StatusCreated {
				var response model.CustomerResponse
				if err := json.NewDecoder(rr.Body).Decode(&response); err != nil {
					t.Fatalf("Could not decode response: %v", err)
				}

				if response.ID != tt.expectedBody.ID {
					t.Errorf("handler returned wrong ID: got %v want %v",
						response.ID, tt.expectedBody.ID)
				}
				if response.Name != tt.expectedBody.Name {
					t.Errorf("handler returned wrong Name: got %v want %v",
						response.Name, tt.expectedBody.Name)
				}
				if (response.EmployerID == nil) != (tt.expectedBody.EmployerID == nil) {
					t.Errorf("handler returned wrong EmployerID presence: got %v want %v",
						response.EmployerID != nil, tt.expectedBody.EmployerID != nil)
				}
				if response.EmployerID != nil && tt.expectedBody.EmployerID != nil {
					if *response.EmployerID != *tt.expectedBody.EmployerID {
						t.Errorf("handler returned wrong EmployerID: got %v want %v",
							*response.EmployerID, *tt.expectedBody.EmployerID)
					}
				}
			}
		})
	}
}

// Helper function to create a pointer to uint
func uintPtr(n uint) *uint {
	return &n
}
