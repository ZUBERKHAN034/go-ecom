package book

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

	router.HandleFunc("/book", handler.book).Methods("POST")
}

func (handler *Handler) book(res http.ResponseWriter, req *http.Request) {}
