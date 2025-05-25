package routes

import (
	"net/http"

	"github.com/ZUBERKHAN034/go-ecom/cmd/controllers"
	"github.com/ZUBERKHAN034/go-ecom/cmd/middlewares"
	"github.com/gorilla/mux"
)

func ProductRoutes(router *mux.Router) {
	router.Handle("/product", middlewares.AuthWithJWT(http.HandlerFunc(controllers.Product.Create))).Methods(http.MethodPost)
	router.Handle("/product/{id}", middlewares.AuthWithJWT(http.HandlerFunc(controllers.Product.Get))).Methods(http.MethodGet)
	router.Handle("/products", middlewares.AuthWithJWT(http.HandlerFunc(controllers.Product.GetAll))).Methods(http.MethodGet)
}
