package routes

import (
	"github.com/ZUBERKHAN034/go-ecom/pkg/routes/book"
	"github.com/ZUBERKHAN034/go-ecom/pkg/routes/user"
	"github.com/gorilla/mux"
)

func RegisterRoutes(router *mux.Router) {
	user.UserRoutes(router)
	book.BookRoutes(router)
}
