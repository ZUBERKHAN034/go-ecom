package controllers

import (
	"net/http"
	"strconv"

	"github.com/ZUBERKHAN034/go-ecom/pkg/lib"
	"github.com/ZUBERKHAN034/go-ecom/pkg/models"
	"github.com/ZUBERKHAN034/go-ecom/pkg/validations"
	"github.com/gorilla/mux"
)

type productController struct{}

type CreateProductPayload struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Image       string  `json:"image"`
	Price       float64 `json:"price"`
	Quantity    int     `json:"quantity"`
}

// CreateProduct godoc
//
// @Summary Create Product
// @Description Create Product
// @Tags Product
// @Accept json
// @Produce json
// @Param payload body CreateProductPayload true "Product Payload"
// @Success 200 {string} string "Product created successfully"
// @Failure 400 {string} string "Invalid request payload"
// @Failure 400 {string} string "Product already exists"
// @Router /product [post]
func (p *productController) CreateProduct(res http.ResponseWriter, req *http.Request) {
	var payload CreateProductPayload

	// Parse the request body
	err := lib.ParseJSON(req, &payload)
	if err != nil {
		lib.SendErrorResponse(res, http.StatusBadRequest, err.Error())
		return
	}

	// validate the payload
	if err := validations.Product.CreateProduct(payload); err != nil {
		lib.SendErrorResponse(res, http.StatusBadRequest, err.Error())
		return
	}

	// check if product already exists
	if checkProduct := models.Product.GetByName(payload.Name); checkProduct.ID != 0 {
		lib.SendErrorResponse(res, http.StatusBadRequest, "Product already exists")
		return
	}

	// create new product
	product := models.Product.Create(&models.ProductSchema{
		Name:        payload.Name,
		Description: payload.Description,
		Image:       payload.Image,
		Price:       payload.Price,
		Quantity:    payload.Quantity,
	})

	lib.SendSuccessResponse(res, http.StatusOK, product)
}

func (p *productController) GetProduct(res http.ResponseWriter, req *http.Request) {
	// getting product by id
	idStr := mux.Vars(req)["id"]
	// convert id string to uint
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		lib.SendErrorResponse(res, http.StatusBadRequest, "Invalid product ID")
		return
	}

	product := models.Product.GetByID(uint(id))
	if product.ID == 0 {
		lib.SendErrorResponse(res, http.StatusNotFound, "Product not found")
		return
	}

	lib.SendSuccessResponse(res, http.StatusOK, product)
}

// GetProducts godoc
//
// @Summary Get Products
// @Description Get Products
// @Tags Product
// @Accept json
// @Produce json
// @Success 200 {array} models.ProductSchema "List of products"
// @Router /products [get]
func (p *productController) GetProducts(res http.ResponseWriter, req *http.Request) {
	// getting products from the database
	products := models.Product.GetAll()
	lib.SendSuccessResponse(res, http.StatusOK, products)
}

var Product = &productController{}
