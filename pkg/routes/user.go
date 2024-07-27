package routes

import (
	"net/http"

	"github.com/ZUBERKHAN034/go-ecom/pkg/controllers"
	"github.com/gorilla/mux"
)

func UserRoutes(router *mux.Router) {
	router.HandleFunc("/user/register", controllers.User.Register).Methods(http.MethodPost)
	router.HandleFunc("/user/login", controllers.User.Login).Methods(http.MethodPost)
}
