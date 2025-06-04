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

	"github.com/gorilla/mux"
)

func TestInvestmentHandler_Create(t *testing.T) {
	tests := []struct {
		name           string
		requestBody    model.InvestmentCreate
		mockInvestment *model.Investment
		mockErr        error
		expectedStatus int
		expectedBody   model.InvestmentResponse
		expectedError  string
	}{
		{
			name: "Create investment successfully",
			requestBody: model.InvestmentCreate{
				ClientID: 1,
				FundID:   1,
				Amount:   1000.0,
			},
			mockInvestment: &model.Investment{
				ID:       1,
				ClientID: 1,
				FundID:   1,
				Amount:   1000.0,
			},
			mockErr:        nil,
			expectedStatus: http.StatusCreated,
			expectedBody: model.InvestmentResponse{
				ID:       1,
				ClientID: 1,
				FundID:   1,
				Amount:   1000.0,
			},
			expectedError: "",
		},
		{
			name: "Empty client ID",
			requestBody: model.InvestmentCreate{
				FundID: 1,
				Amount: 1000.0,
			},
			mockInvestment: nil,
			mockErr:        nil,
			expectedStatus: http.StatusBadRequest,
			expectedBody:   model.InvestmentResponse{},
			expectedError:  "Client ID is required",
		},
		{
			name: "Empty fund ID",
			requestBody: model.InvestmentCreate{
				ClientID: 1,
				Amount:   1000.0,
			},
			mockInvestment: nil,
			mockErr:        nil,
			expectedStatus: http.StatusBadRequest,
			expectedBody:   model.InvestmentResponse{},
			expectedError:  "Fund ID is required",
		},
		{
			name: "Zero amount",
			requestBody: model.InvestmentCreate{
				ClientID: 1,
				FundID:   1,
				Amount:   0,
			},
			mockInvestment: nil,
			mockErr:        nil,
			expectedStatus: http.StatusBadRequest,
			expectedBody:   model.InvestmentResponse{},
			expectedError:  "Amount must be greater than 0",
		},
		{
			name: "Negative amount",
			requestBody: model.InvestmentCreate{
				ClientID: 1,
				FundID:   1,
				Amount:   -1000.0,
			},
			mockInvestment: nil,
			mockErr:        nil,
			expectedStatus: http.StatusBadRequest,
			expectedBody:   model.InvestmentResponse{},
			expectedError:  "Amount must be greater than 0",
		},
		{
			name:           "Invalid request body",
			requestBody:    model.InvestmentCreate{},
			mockInvestment: nil,
			mockErr:        nil,
			expectedStatus: http.StatusBadRequest,
			expectedBody:   model.InvestmentResponse{},
			expectedError:  "Invalid request body",
		},
		{
			name: "Service error",
			requestBody: model.InvestmentCreate{
				ClientID: 1,
				FundID:   1,
				Amount:   1000.0,
			},
			mockInvestment: nil,
			mockErr:        errors.New("service error"),
			expectedStatus: http.StatusBadRequest,
			expectedBody:   model.InvestmentResponse{},
			expectedError:  "service error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockService := &mocks.InvestmentService{
				MockInvestment: tt.mockInvestment,
				MockErr:        tt.mockErr,
			}

			handler := NewInvestmentHandler(mockService)

			var req *http.Request
			if tt.name == "Invalid request body" {
				req = httptest.NewRequest("POST", "/investments", bytes.NewBufferString("invalid json"))
			} else {
				body, _ := json.Marshal(tt.requestBody)
				req = httptest.NewRequest("POST", "/investments", bytes.NewBuffer(body))
			}

			rr := httptest.NewRecorder()

			handler.Create(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v",
					rr.Code, tt.expectedStatus)
			}

			if tt.expectedStatus == http.StatusCreated {
				var response model.InvestmentResponse
				if err := json.NewDecoder(rr.Body).Decode(&response); err != nil {
					t.Fatalf("Could not decode response: %v", err)
				}

				if response.ID != tt.expectedBody.ID {
					t.Errorf("handler returned wrong ID: got %v want %v",
						response.ID, tt.expectedBody.ID)
				}
				if response.ClientID != tt.expectedBody.ClientID {
					t.Errorf("handler returned wrong ClientID: got %v want %v",
						response.ClientID, tt.expectedBody.ClientID)
				}
				if response.FundID != tt.expectedBody.FundID {
					t.Errorf("handler returned wrong FundID: got %v want %v",
						response.FundID, tt.expectedBody.FundID)
				}
				if response.Amount != tt.expectedBody.Amount {
					t.Errorf("handler returned wrong Amount: got %v want %v",
						response.Amount, tt.expectedBody.Amount)
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

func TestInvestmentHandler_Get(t *testing.T) {
	tests := []struct {
		name           string
		investmentID   string
		mockInvestment *model.Investment
		mockErr        error
		expectedStatus int
		expectedBody   model.InvestmentResponse
		expectedError  string
	}{
		{
			name:         "Get investment successfully",
			investmentID: "1",
			mockInvestment: &model.Investment{
				ID:       1,
				ClientID: 1,
				FundID:   1,
				Amount:   1000.0,
			},
			mockErr:        nil,
			expectedStatus: http.StatusOK,
			expectedBody: model.InvestmentResponse{
				ID:       1,
				ClientID: 1,
				FundID:   1,
				Amount:   1000.0,
			},
			expectedError: "",
		},
		{
			name:           "Invalid investment ID",
			investmentID:   "invalid",
			mockInvestment: nil,
			mockErr:        nil,
			expectedStatus: http.StatusBadRequest,
			expectedBody:   model.InvestmentResponse{},
			expectedError:  "Invalid investment ID",
		},
		{
			name:           "Investment not found",
			investmentID:   "999",
			mockInvestment: nil,
			mockErr:        errors.New("investment not found"),
			expectedStatus: http.StatusNotFound,
			expectedBody:   model.InvestmentResponse{},
			expectedError:  "investment not found",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockService := &mocks.InvestmentService{
				MockInvestment: tt.mockInvestment,
				MockErr:        tt.mockErr,
			}

			handler := NewInvestmentHandler(mockService)

			req := httptest.NewRequest("GET", "/investments/"+tt.investmentID, nil)
			rr := httptest.NewRecorder()

			// Create a new router and register the handler
			router := mux.NewRouter()
			router.HandleFunc("/investments/{id}", handler.Get).Methods("GET")

			// Serve the request
			router.ServeHTTP(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v",
					rr.Code, tt.expectedStatus)
			}

			if tt.expectedStatus == http.StatusOK {
				var response model.InvestmentResponse
				if err := json.NewDecoder(rr.Body).Decode(&response); err != nil {
					t.Fatalf("Could not decode response: %v", err)
				}

				if response.ID != tt.expectedBody.ID {
					t.Errorf("handler returned wrong ID: got %v want %v",
						response.ID, tt.expectedBody.ID)
				}
				if response.ClientID != tt.expectedBody.ClientID {
					t.Errorf("handler returned wrong ClientID: got %v want %v",
						response.ClientID, tt.expectedBody.ClientID)
				}
				if response.FundID != tt.expectedBody.FundID {
					t.Errorf("handler returned wrong FundID: got %v want %v",
						response.FundID, tt.expectedBody.FundID)
				}
				if response.Amount != tt.expectedBody.Amount {
					t.Errorf("handler returned wrong Amount: got %v want %v",
						response.Amount, tt.expectedBody.Amount)
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

func TestInvestmentHandler_GetAll(t *testing.T) {
	tests := []struct {
		name            string
		clientID        string
		mockInvestments []*model.Investment
		mockErr         error
		expectedStatus  int
		expectedBody    []model.InvestmentResponse
		expectedError   string
	}{
		{
			name:     "Get investments successfully",
			clientID: "1",
			mockInvestments: []*model.Investment{
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
			mockErr:        nil,
			expectedStatus: http.StatusOK,
			expectedBody: []model.InvestmentResponse{
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
			expectedError: "",
		},
		{
			name:            "No investments found",
			clientID:        "1",
			mockInvestments: []*model.Investment{},
			mockErr:         nil,
			expectedStatus:  http.StatusOK,
			expectedBody:    []model.InvestmentResponse{},
			expectedError:   "",
		},
		{
			name:            "Invalid client ID",
			clientID:        "invalid",
			mockInvestments: nil,
			mockErr:         nil,
			expectedStatus:  http.StatusBadRequest,
			expectedBody:    nil,
			expectedError:   "Invalid client ID",
		},
		{
			name:            "Service error",
			clientID:        "1",
			mockInvestments: nil,
			mockErr:         errors.New("service error"),
			expectedStatus:  http.StatusInternalServerError,
			expectedBody:    nil,
			expectedError:   "service error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockService := &mocks.InvestmentService{
				MockInvestments: tt.mockInvestments,
				MockErr:         tt.mockErr,
			}

			handler := NewInvestmentHandler(mockService)

			req := httptest.NewRequest("GET", "/investments?client_id="+tt.clientID, nil)
			rr := httptest.NewRecorder()

			handler.GetAll(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v",
					rr.Code, tt.expectedStatus)
			}

			if tt.expectedStatus == http.StatusOK {
				var response []model.InvestmentResponse
				if err := json.NewDecoder(rr.Body).Decode(&response); err != nil {
					t.Fatalf("Could not decode response: %v", err)
				}

				if len(response) != len(tt.expectedBody) {
					t.Errorf("handler returned wrong number of investments: got %v want %v",
						len(response), len(tt.expectedBody))
					return
				}

				for i, expected := range tt.expectedBody {
					if response[i].ID != expected.ID {
						t.Errorf("investment[%d].ID = %v, want %v", i, response[i].ID, expected.ID)
					}
					if response[i].ClientID != expected.ClientID {
						t.Errorf("investment[%d].ClientID = %v, want %v", i, response[i].ClientID, expected.ClientID)
					}
					if response[i].FundID != expected.FundID {
						t.Errorf("investment[%d].FundID = %v, want %v", i, response[i].FundID, expected.FundID)
					}
					if response[i].Amount != expected.Amount {
						t.Errorf("investment[%d].Amount = %v, want %v", i, response[i].Amount, expected.Amount)
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
