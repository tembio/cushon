package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"cushon/internal/mocks"
	"cushon/internal/model"
)

func TestEmployerHandler_Create(t *testing.T) {
	tests := []struct {
		name           string
		requestBody    model.EmployerCreate
		mockEmployer   *model.Employer
		mockErr        error
		expectedStatus int
		expectedBody   model.EmployerResponse
		expectedError  string
	}{
		{
			name: "Create employer successfully",
			requestBody: model.EmployerCreate{
				Name: "Test Company",
			},
			mockEmployer: &model.Employer{
				ID:   1,
				Name: "Test Company",
			},
			mockErr:        nil,
			expectedStatus: http.StatusCreated,
			expectedBody: model.EmployerResponse{
				ID:   1,
				Name: "Test Company",
			},
			expectedError: "",
		},
		{
			name: "Invalid request body",
			requestBody: model.EmployerCreate{
				Name: "Invalid JSON",
			},
			mockEmployer:   nil,
			mockErr:        nil,
			expectedStatus: http.StatusBadRequest,
			expectedBody:   model.EmployerResponse{},
			expectedError:  "Invalid request body",
		},
		{
			name: "Service error",
			requestBody: model.EmployerCreate{
				Name: "Test Company",
			},
			mockEmployer:   nil,
			mockErr:        errors.New("service error"),
			expectedStatus: http.StatusBadRequest,
			expectedBody:   model.EmployerResponse{},
			expectedError:  "service error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockService := &mocks.EmployerService{
				MockEmployer: tt.mockEmployer,
				MockErr:      tt.mockErr,
			}

			handler := NewEmployerHandler(mockService)

			var req *http.Request
			if tt.name == "Invalid request body" {
				req = httptest.NewRequest("POST", "/employers", bytes.NewBufferString("invalid json"))
			} else {
				body, _ := json.Marshal(tt.requestBody)
				req = httptest.NewRequest("POST", "/employers", bytes.NewBuffer(body))
			}

			rr := httptest.NewRecorder()

			handler.Create(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v",
					rr.Code, tt.expectedStatus)
			}

			if tt.expectedStatus == http.StatusCreated {
				var response model.EmployerResponse
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
			} else if tt.expectedError != "" {
				if rr.Body.String() != tt.expectedError+"\n" {
					t.Errorf("handler returned wrong error message: got %v want %v",
						rr.Body.String(), tt.expectedError)
				}
			}
		})
	}
}
