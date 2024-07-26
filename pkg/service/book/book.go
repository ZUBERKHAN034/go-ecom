package book

import (
	"fmt"
	"net/http"
)

func Book(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Book service")
	res.Write([]byte("Hello Book Service"))
}
