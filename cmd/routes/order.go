package routes

import (
	"net/http"

	"github.com/ZUBERKHAN034/go-ecom/cmd/controllers"
	"github.com/ZUBERKHAN034/go-ecom/cmd/middlewares"
	"github.com/gorilla/mux"
)

func OrderRoutes(router *mux.Router) {
	router.Handle("/order/checkout", middlewares.AuthWithJWT(http.HandlerFunc(controllers.Order.Checkout))).Methods(http.MethodPost)
}
