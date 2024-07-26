package user

import (
	"github.com/ZUBERKHAN034/go-ecom/pkg/service/user"
	"github.com/gorilla/mux"
)

func UserRoutes(router *mux.Router) {
	router.HandleFunc("/register", user.Register).Methods("POST")
	router.HandleFunc("/login", user.Login).Methods("POST")
}
