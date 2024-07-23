package user

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (handler *Handler) RegisterRoutes(router *mux.Router) {

	router.HandleFunc("/login", handler.login).Methods("POST")
	router.HandleFunc("/register", handler.register).Methods("POST")
}

func (handler *Handler) login(res http.ResponseWriter, req *http.Request)    {}
func (handler *Handler) register(res http.ResponseWriter, req *http.Request) {}
