package routes

import (
	"github.com/ZUBERKHAN034/go-ecom/pkg/services"
	"github.com/gorilla/mux"
)

func UserRoutes(router *mux.Router) {
	router.HandleFunc("/user/register", services.User.Register).Methods("POST")
	router.HandleFunc("/user/login", services.User.Login).Methods("POST")
}
