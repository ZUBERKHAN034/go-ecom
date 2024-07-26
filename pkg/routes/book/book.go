package book

import (
	"github.com/ZUBERKHAN034/go-ecom/pkg/service/book"
	"github.com/gorilla/mux"
)

func BookRoutes(router *mux.Router) {
	router.HandleFunc("/book", book.Book).Methods("GET")
}
