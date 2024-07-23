package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/ZUBERKHAN034/go-ecom/service/user"
	"github.com/gorilla/mux"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (apiServer *APIServer) RUN() error {
	pathPrefix := "/api/v1"

	router := mux.NewRouter()
	subRouter := router.PathPrefix(pathPrefix).Subrouter()

	userHandler := user.NewHandler()
	userHandler.RegisterRoutes(subRouter)

	log.Println("Listening on", apiServer.addr)
	return http.ListenAndServe(apiServer.addr, router)
}
