package controllers

import (
	"errors"
	"net/http"

	"github.com/ZUBERKHAN034/go-ecom/pkg/lib"
	"github.com/ZUBERKHAN034/go-ecom/pkg/models"
)

type bookController struct{}

type CreateBookPayload struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Pages  int    `json:"pages"`
	Price  int    `json:"price"`
}

//	CreateBook godoc
//
// @Summary Create Book
// @Description Create Book
// @Tags Book
// @Accept json
// @Produce json
// @Param payload body CreateBookPayload true "Book Payload"
// @Success 200 {string} string "Book created successfully"
// @Failure 400 {string} string "Invalid request payload"
// @Failure 400 {string} string "Book already exists"
// @Router /book [post]
func (b *bookController) CreateBook(res http.ResponseWriter, req *http.Request) {
	var payload CreateBookPayload

	// Parse the request body
	err := lib.ParseJSON(req, &payload)
	if err != nil {
		lib.SendErrorResponse(res, http.StatusBadRequest, err)
		return
	}

	// check if book already exists
	book := models.Book.GetByTitle(payload.Title)
	if book.ID != 0 {
		lib.SendErrorResponse(res, http.StatusBadRequest, errors.New("book already exists"))
		return
	}

	// create new book
	models.Book.Create(&models.BookSchema{
		Title:  payload.Title,
		Author: payload.Author,
		Pages:  payload.Pages,
		Price:  payload.Price,
	})

	lib.SendErrorResponse(res, http.StatusOK, "book created successfully")
}

//	GetBook godoc
//
// @Summary Get Books
// @Description Get Book
// @Tags Book
// @Accept json
// @Produce json
// @Success 200 {string} string "Hello Book Service"
// @Router /books [get]
func (b *bookController) GetBooks(res http.ResponseWriter, req *http.Request) {
	books := models.Book.GetAll()
	lib.SendSuccessResponse(res, http.StatusOK, books)
}

var Book = &bookController{}
