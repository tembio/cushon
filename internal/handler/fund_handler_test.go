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

func TestFundHandler_Create(t *testing.T) {
	tests := []struct {
		name           string
		requestBody    model.FundCreate
		mockFund       *model.Fund
		mockErr        error
		expectedStatus int
		expectedBody   model.FundResponse
		expectedError  string
	}{
		{
			name: "Create fund successfully",
			requestBody: model.FundCreate{
				Name: "Test Fund",
			},
			mockFund: &model.Fund{
				ID:   1,
				Name: "Test Fund",
			},
			mockErr:        nil,
			expectedStatus: http.StatusCreated,
			expectedBody: model.FundResponse{
				ID:   1,
				Name: "Test Fund",
			},
			expectedError: "",
		},
		{
			name:           "Invalid request body",
			requestBody:    model.FundCreate{},
			mockFund:       nil,
			mockErr:        nil,
			expectedStatus: http.StatusBadRequest,
			expectedBody:   model.FundResponse{},
			expectedError:  "Invalid request body",
		},
		{
			name:           "Service error",
			requestBody:    model.FundCreate{},
			mockFund:       nil,
			mockErr:        errors.New("service error"),
			expectedStatus: http.StatusBadRequest,
			expectedBody:   model.FundResponse{},
			expectedError:  "service error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockService := &mocks.FundService{
				MockFund: tt.mockFund,
				MockErr:  tt.mockErr,
			}

			handler := NewFundHandler(mockService)

			var req *http.Request
			if tt.name == "Invalid request body" {
				req = httptest.NewRequest("POST", "/funds", bytes.NewBufferString("invalid json"))
			} else {
				body, _ := json.Marshal(tt.requestBody)
				req = httptest.NewRequest("POST", "/funds", bytes.NewBuffer(body))
			}

			rr := httptest.NewRecorder()

			handler.Create(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v",
					rr.Code, tt.expectedStatus)
			}

			if tt.expectedStatus == http.StatusCreated {
				var response model.FundResponse
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

func TestFundHandler_GetAll(t *testing.T) {
	tests := []struct {
		name           string
		mockFunds      []*model.Fund
		mockErr        error
		expectedStatus int
		expectedBody   []model.FundResponse
		expectedError  string
	}{
		{
			name: "Get all funds successfully",
			mockFunds: []*model.Fund{
				{
					ID:   1,
					Name: "Test Fund 1",
				},
				{
					ID:   2,
					Name: "Test Fund 2",
				},
			},
			mockErr:        nil,
			expectedStatus: http.StatusOK,
			expectedBody: []model.FundResponse{
				{
					ID:   1,
					Name: "Test Fund 1",
				},
				{
					ID:   2,
					Name: "Test Fund 2",
				},
			},
			expectedError: "",
		},
		{
			name:           "Empty funds list",
			mockFunds:      []*model.Fund{},
			mockErr:        nil,
			expectedStatus: http.StatusOK,
			expectedBody:   []model.FundResponse{},
			expectedError:  "",
		},
		{
			name:           "Service error",
			mockFunds:      nil,
			mockErr:        errors.New("service error"),
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   nil,
			expectedError:  "service error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockService := &mocks.FundService{
				MockFunds: tt.mockFunds,
				MockErr:   tt.mockErr,
			}

			handler := NewFundHandler(mockService)
			req := httptest.NewRequest("GET", "/funds", nil)
			rr := httptest.NewRecorder()

			handler.GetAll(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v",
					rr.Code, tt.expectedStatus)
			}

			if tt.expectedStatus == http.StatusOK {
				var response []model.FundResponse
				if err := json.NewDecoder(rr.Body).Decode(&response); err != nil {
					t.Fatalf("Could not decode response: %v", err)
				}

				if len(response) != len(tt.expectedBody) {
					t.Errorf("handler returned wrong number of funds: got %v want %v",
						len(response), len(tt.expectedBody))
					return
				}

				for i, fund := range response {
					if fund.ID != tt.expectedBody[i].ID {
						t.Errorf("handler returned wrong ID for fund %d: got %v want %v",
							i, fund.ID, tt.expectedBody[i].ID)
					}
					if fund.Name != tt.expectedBody[i].Name {
						t.Errorf("handler returned wrong Name for fund %d: got %v want %v",
							i, fund.Name, tt.expectedBody[i].Name)
					}
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
