package routes

import (
	"net/http"

	"github.com/ZUBERKHAN034/go-ecom/pkg/controllers"
	"github.com/gorilla/mux"
)

func BookRoutes(router *mux.Router) {
	router.HandleFunc("/book", controllers.Book.GetBook).Methods(http.MethodGet)
}
