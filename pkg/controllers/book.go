package controllers

import (
	"net/http"
)

type bookController struct{}

//	GetBook godoc
//
// @Summary Get Book
// @Description Get Book
// @Tags Book
// @Accept json
// @Produce json
// @Success 200 {string} string "Hello Book Service"
// @Router /book [get]
func (b *bookController) GetBook(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Hello Book Service"))
}

var Book = &bookController{}
