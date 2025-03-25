package controllers

import (
	"fmt"
	"net/http"

	"github.com/ZUBERKHAN034/go-ecom/pkg/app"
	"github.com/ZUBERKHAN034/go-ecom/pkg/models"
	"github.com/ZUBERKHAN034/go-ecom/pkg/validations"
)

type orderController struct{}

// Checkout godoc
// 
// @Summary Checkout
// @Description Checkout
// @Tags Order
// @Accept json
// @Produce json
// @Param payload body types.OrderPayload true "Order Payload"
// @Success 200 {object} models.OrderSchema "Order"
// @Failure 400 {string} string "invalid request payload"
// @Failure 400 {string} string "product not found for ID: {productId}"
// @Failure 400 {string} string "order items can not be empty"
// @Failure 400 {string} string "invalid quantity for product ID: {productId}"
// @Failure 400 {string} string "invalid product IDs"
// @Failure 500 {string} string "internal server error"
// @Router /order/checkout [post]
func (o *orderController) Checkout(res http.ResponseWriter, req *http.Request) {

	// validate the request payload
	order, err := validations.Order.Checkout(req)
	if err != nil {
		app.SendErrorResponse(res, http.StatusBadRequest, err.Error())
		return
	}

	// check if products exist in the database
	for _, item := range order.Items {
		product := models.Product.GetByID(item.ProductID)
		if product.ID == 0 {
			errMsg := fmt.Sprintf("product not found for ID: %d", item.ProductID)
			app.SendErrorResponse(res, http.StatusBadRequest, errMsg)
			return
		}
	}

	// check if order items are empty
	if len(order.Items) == 0 {
		app.SendErrorResponse(res, http.StatusBadRequest, "order items can not be empty")
		return
	}

	// validate the order Items and set the Items IDS in the productIDs slice
	productIDs := make([]uint, len(order.Items))
	for i, item := range order.Items {
		if item.Quantity <= 0 {
			app.SendErrorResponse(res, http.StatusBadRequest, fmt.Errorf("invalid quantity for product ID: %d", item.ProductID))
			return
		}

		productIDs[i] = item.ProductID
	}

	// get products from the database
	products := models.Product.GetProductsByIDs(productIDs)
	if len(products) != len(order.Items) {
		app.SendErrorResponse(res, http.StatusBadRequest, "invalid product IDs")
		return
	}

	app.SendSuccessResponse(res, http.StatusOK, "order processed successfully")
}

var Order = &orderController{}
