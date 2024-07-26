package main

import (
	"log"

	"github.com/ZUBERKHAN034/go-ecom/cmd/app"
	"github.com/ZUBERKHAN034/go-ecom/pkg/config"
	"github.com/ZUBERKHAN034/go-ecom/pkg/db"
)

func main() {
	addr := ":8080"
	if config.Env.Port != "" {
		addr = ":" + config.Env.Port
	}

	db, err := db.ConnectMySQL()
	if err != nil {
		log.Fatal(err)
	}

	server := app.InitAPIServer(addr, db)
	if err := server.RUN(); err != nil {
		log.Fatal(err)
	}
}
