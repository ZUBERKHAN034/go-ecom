package main

import (
	"log"

	"github.com/ZUBERKHAN034/go-ecom/cmd/api"
	"github.com/ZUBERKHAN034/go-ecom/db"
)

func main() {
	addr := ":8080"

	db, err := db.ConnectMySQL()
	if err != nil {
		log.Fatal(err)
	}

	server := api.InitAPIServer(addr, db)
	if err := server.RUN(); err != nil {
		log.Fatal(err)
	}
}
