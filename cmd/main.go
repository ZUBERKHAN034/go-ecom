package main

import (
	"log"
	"net/http"

	"github.com/ZUBERKHAN034/go-ecom/cmd/config"
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
	addr := ":8080"
	if config.Env.Port != "" {
		addr = ":" + config.Env.Port
	}

	// Create router
	router := mux.NewRouter()

	// Initialize routes groups
	routes.InitRoutes(router)

	// Start the server
	if err := http.ListenAndServe(addr, router); err != nil {
		log.Fatal("Server failed:", err)
	} else {
		log.Println("Listening on", addr)
	}
}
