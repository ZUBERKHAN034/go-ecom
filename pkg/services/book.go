package services

import (
	"fmt"
	"net/http"
)

type bookService struct{}

func (bs *bookService) GetBook(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Book service")
	res.Write([]byte("Hello Book Service"))
}

var Book = &bookService{}
