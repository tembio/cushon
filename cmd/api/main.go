package main

import (
	"log"
	"net/http"

	"cushon/internal/app/handler"
	"cushon/internal/app/middleware"
	"cushon/internal/app/repository"
	"cushon/internal/app/service"

	"github.com/gorilla/mux"
)

func main() {
	fundRepo := repository.NewInMemoryFundRepository()
	apiKeyRepo := repository.NewInMemoryAPIKeyRepository()
	fundService := service.NewDefaultFundService(fundRepo)
	fundHandler := handler.NewFundHandler(fundService)
	router := mux.NewRouter()

	// Protected routes
	routes := router.PathPrefix("/").Subrouter()
	routes.Use(middleware.NewAuthMiddleware(apiKeyRepo))
	routes.HandleFunc("/funds", fundHandler.Create).Methods("POST")

	log.Println("Starting server on :8443")
	err := http.ListenAndServeTLS(":8443", "certs/server.crt", "certs/server.key", router)
	if err != nil {
		log.Fatal(err)
	}
}
