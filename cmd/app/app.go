package app

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/ZUBERKHAN034/go-ecom/pkg/routes/user"
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
	pathPrefix := "/api/v1"

	router := mux.NewRouter()
	user.UserRoutes(router)

	http.Handle(pathPrefix, router)

	log.Println("Listening on", apiServer.addr)
	return http.ListenAndServe(apiServer.addr, router)
}
