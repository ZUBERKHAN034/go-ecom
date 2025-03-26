package controllers

import (
	"net/http"
	"strconv"

	"github.com/ZUBERKHAN034/go-ecom/pkg/app"
	"github.com/ZUBERKHAN034/go-ecom/pkg/models"
	"github.com/ZUBERKHAN034/go-ecom/pkg/validations"
	"github.com/gorilla/mux"
)

type productController struct{}

// Create godoc
//
// @Summary Create Product
// @Description Create Product
// @Tags Product
// @Accept json
// @Produce json
// @Param payload body types.ProductPayload true "Product Payload"
// @Success 201 {object} models.ProductSchema "Product"
// @Failure 400 {string} string "invalid request payload"
// @Failure 400 {string} string "product already exists"
// @Failure 500 {string} string "internal server error"
// @Router /product [post]
func (p *productController) Create(res http.ResponseWriter, req *http.Request) {

	// validate the request payload
	product, err := validations.Product.Create(req)
	if err != nil {
		app.SendErrorResponse(res, http.StatusBadRequest, err.Error())
		return
	}

	// check if product already exists
	if checkProduct := models.Product.GetByName(product.Name); checkProduct.ID != 0 {
		app.SendErrorResponse(res, http.StatusBadRequest, "product already exists")
		return
	}

	// create new product
	createdProduct := models.Product.Create(&models.ProductSchema{
		Name:        product.Name,
		Description: product.Description,
		Image:       product.Image,
		Price:       product.Price,
		Quantity:    product.Quantity,
	})

	app.SendSuccessResponse(res, http.StatusCreated, createdProduct)
}

// Get godoc
//
// @Summary Get Product
// @Description Get Product
// @Tags Product
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Success 200 {object} models.ProductSchema "Product"
// @Failure 400 {string} string "invalid product ID"
// @Failure 404 {string} string "product not exists"
// @Failure 500 {string} string "internal server error"
// @Router /product/{id} [get]
func (p *productController) Get(res http.ResponseWriter, req *http.Request) {

	// get product by id
	idStr := mux.Vars(req)["id"]
	// convert id string to uint
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		app.SendErrorResponse(res, http.StatusBadRequest, "invalid product ID")
		return
	}

	product := models.Product.GetByID(uint(id))
	if product.ID == 0 {
		app.SendErrorResponse(res, http.StatusNotFound, "product not exists")
		return
	}

	app.SendSuccessResponse(res, http.StatusOK, product)
}

// GetAll godoc
//
// @Summary Get Products
// @Description Get Products
// @Tags Product
// @Accept json
// @Produce json
// @Success 200 {array} models.ProductSchema "List of products"
// @Failure 500 {string} string "internal server error"
// @Router /products [get]
func (p *productController) GetAll(res http.ResponseWriter, req *http.Request) {

	// get the products from the database
	products := models.Product.GetAll()
	app.SendSuccessResponse(res, http.StatusOK, products)
}

var Product = &productController{}
