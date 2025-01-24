package controllers

import (
	"fmt"
	"net/http"

	"github.com/ZUBERKHAN034/go-ecom/pkg/lib"
	"github.com/ZUBERKHAN034/go-ecom/pkg/models"
	"github.com/ZUBERKHAN034/go-ecom/pkg/validations"
)

type orderController struct{}

func (o *orderController) Checkout(res http.ResponseWriter, req *http.Request) {
	fmt.Println("CHECKOUT CALLED")

	// validate the payload
	order, err := validations.Order.Checkout(req)
	if err != nil {
		lib.SendErrorResponse(res, http.StatusBadRequest, err.Error())
		return
	}

	// parse request body
	// if err := lib.ParseJSON(req, &order); err != nil {
	// 	lib.SendErrorResponse(res, http.StatusBadRequest, err.Error())
	// 	return
	// }

	// check if order items are empty
	if len(order.Items) == 0 {
		lib.SendErrorResponse(res, http.StatusBadRequest, "Order items can not be empty")
		return
	}

	// validate the order Items and set the Items IDS in the productIDs slice
	productIDs := make([]uint, len(order.Items))
	for i, item := range order.Items {
		if item.Quantity <= 0 {
			lib.SendErrorResponse(res, http.StatusBadRequest, fmt.Errorf("invalid quantity for product ID: %d", item.ProductID))
			return
		}

		productIDs[i] = item.ProductID
	}

	// get products from the database
	products := models.Product.GetProductsByIDs(productIDs)
	if len(products) != len(order.Items) {
		lib.SendErrorResponse(res, http.StatusBadRequest, "Invalid product IDs")
		return
	}

	lib.SendSuccessResponse(res, http.StatusOK, "Order processed successfully")
}

var Order = &orderController{}
