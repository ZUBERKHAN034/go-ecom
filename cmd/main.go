package main

import (
	"log"
	"net/http"
	"os"

	"github.com/ZUBERKHAN034/go-ecom/cmd/routes"
	"github.com/gorilla/mux"
)

// @title           E-Commerce API
// @version         1.0
// @description     This is E-Commerce API server
// @termsOfService  http://github.com/ZUBERKHAN034/go-ecom
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	addr := ":" + port
	// Create router
	router := mux.NewRouter()

	// Initialize routes groups
	routes.InitRoutes(router)

	log.Println("Starting server on", addr)
	if err := http.ListenAndServe(addr, router); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
