package routes

import (
	"net/http"

	"github.com/ZUBERKHAN034/go-ecom/cmd/controllers"
	"github.com/gorilla/mux"
)

func UserRoutes(router *mux.Router) {
	router.HandleFunc("/users/register", controllers.User.Register).Methods(http.MethodPost)
	router.HandleFunc("/users/login", controllers.User.Login).Methods(http.MethodPost)
}
