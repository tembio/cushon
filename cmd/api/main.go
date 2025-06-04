package main

import (
	"log"
	"net/http"
	"os"

	"cushon/internal/handler"
	"cushon/internal/middleware"
	"cushon/internal/repository"
	"cushon/internal/service"

	"github.com/gorilla/mux"
)

func main() {
	// Check if certificates exist
	certPath := "certs/server.crt"
	keyPath := "certs/server.key"
	if _, err := os.Stat(certPath); os.IsNotExist(err) {
		log.Fatal("SSL certificate not found at:", certPath)
	}
	if _, err := os.Stat(keyPath); os.IsNotExist(err) {
		log.Fatal("SSL key not found at:", keyPath)
	}

	// Initialize repositories
	customerRepo := repository.NewInMemoryCustomerRepository()
	fundRepo := repository.NewInMemoryFundRepository()
	investmentRepo := repository.NewInMemoryInvestmentRepository()
	employerRepo := repository.NewInMemoryEmployerRepository()
	apiKeyRepo := repository.NewInMemoryAPIKeyRepository()

	// Add a test API key
	apiKeyRepo.AddKey("test-api-key")

	// Initialize services
	customerService := service.NewDefaultCustomerService(customerRepo)
	fundService := service.NewDefaultFundService(fundRepo)
	investmentService := service.NewDefaultInvestmentService(investmentRepo)
	employerService := service.NewDefaultEmployerService(employerRepo)

	// Initialize handlers
	customerHandler := handler.NewCustomerHandler(customerService)
	fundHandler := handler.NewFundHandler(fundService)
	investmentHandler := handler.NewInvestmentHandler(investmentService)
	employerHandler := handler.NewEmployerHandler(employerService)

	// Create router
	router := mux.NewRouter()

	// Add health check endpoint (no auth required)
	router.HandleFunc("/health", handler.HealthCheck).Methods("GET")

	// Create authenticated subrouter for all other endpoints
	api := router.PathPrefix("/api").Subrouter()
	api.Use(middleware.NewAuthMiddleware(apiKeyRepo))

	// Customer routes
	api.HandleFunc("/customers", customerHandler.Create).Methods("POST")

	// Fund routes
	api.HandleFunc("/funds", fundHandler.Create).Methods("POST")
	api.HandleFunc("/funds", fundHandler.GetAll).Methods("GET")

	// Investment routes
	api.HandleFunc("/investments", investmentHandler.Create).Methods("POST")
	api.HandleFunc("/investments/{id}", investmentHandler.Get).Methods("GET")
	api.HandleFunc("/investments", investmentHandler.GetAll).Methods("GET")

	// Employer routes
	api.HandleFunc("/employers", employerHandler.Create).Methods("POST")

	// Start server
	log.Println("Starting server on :8443")
	err := http.ListenAndServeTLS(":8443", certPath, keyPath, router)
	if err != nil {
		log.Fatal(err)
	}
}
