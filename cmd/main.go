package main

import (
	"log"

	"github.com/ZUBERKHAN034/go-ecom/cmd/app"
	"github.com/ZUBERKHAN034/go-ecom/pkg/config"
)

// @title			E-Commerce API
// @version		1.0
// @description	This is E-Commerce API server
// @termsOfService	http://github.com/ZUBERKHAN034/go-ecom
func main() {
	addr := ":8080"
	if config.Env.Port != "" {
		addr = ":" + config.Env.Port
	}

	server := app.InitAPIServer(addr)
	if err := server.RUN(); err != nil {
		log.Fatal(err)
	}
}
