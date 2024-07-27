package routes

import (
	"github.com/gorilla/mux"
)

func InitRoutes(router *mux.Router) {
	UserRoutes(router)
	BookRoutes(router)
}
