package routes

import (
	"github.com/gorilla/mux"
)

func InitRoutes(router *mux.Router) {
	SwaggerRoutes(router)
	UserRoutes(router)
	ProductRoutes(router)
}
