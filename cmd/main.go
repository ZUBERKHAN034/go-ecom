package main

import (
	"log"

	"github.com/ZUBERKHAN034/go-ecom/cmd/api"
)

func main() {
	addr := ":8080"
	server := api.NewAPIServer(addr, nil)
	if err := server.RUN(); err != nil {
		log.Fatal(err)
	}
}
