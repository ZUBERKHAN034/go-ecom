package app

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/ZUBERKHAN034/go-ecom/pkg/routes"
	"github.com/gorilla/mux"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func InitAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (apiServer *APIServer) RUN() error {
	router := mux.NewRouter()
	routes.RegisterRoutes(router)
	http.Handle("/", router)

	log.Println("Listening on", apiServer.addr)
	return http.ListenAndServe(apiServer.addr, router)
}
