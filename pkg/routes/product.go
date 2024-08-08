package routes

import (
	"net/http"

	"github.com/ZUBERKHAN034/go-ecom/pkg/controllers"
	"github.com/gorilla/mux"
)

func ProductRoutes(router *mux.Router) {
	router.HandleFunc("/product", controllers.Product.CreateProduct).Methods(http.MethodPost)
	router.HandleFunc("/product/{id}", controllers.Product.GetProduct).Methods(http.MethodGet)
	router.HandleFunc("/products", controllers.Product.GetProducts).Methods(http.MethodGet)
}
