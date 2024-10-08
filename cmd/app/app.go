package app

import (
	"log"
	"net/http"

	"github.com/ZUBERKHAN034/go-ecom/pkg/routes"
	"github.com/gorilla/mux"
)

type APIServer struct {
	addr string
}

func InitAPIServer(addr string) *APIServer {
	return &APIServer{
		addr: addr,
	}
}

func (apiServer *APIServer) RUN() error {
	router := mux.NewRouter()
	routes.InitRoutes(router)

	log.Println("Listening on", apiServer.addr)
	log.Println("Swagger UI is available at http://localhost:8080/swagger/index.html")

	return http.ListenAndServe(apiServer.addr, router)
}
