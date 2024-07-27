package routes

import (
	"github.com/ZUBERKHAN034/go-ecom/pkg/services"
	"github.com/gorilla/mux"
)

func BookRoutes(router *mux.Router) {
	router.HandleFunc("/book", services.Book.GetBook).Methods("GET")
}
