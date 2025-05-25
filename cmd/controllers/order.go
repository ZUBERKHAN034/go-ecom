package controllers

import (
	"fmt"
	"net/http"

	"github.com/ZUBERKHAN034/go-ecom/cmd/app"
	"github.com/ZUBERKHAN034/go-ecom/cmd/middlewares"
	"github.com/ZUBERKHAN034/go-ecom/cmd/models"
	"github.com/ZUBERKHAN034/go-ecom/cmd/types"
	"github.com/ZUBERKHAN034/go-ecom/cmd/validations"
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
// @Success 201 {object} types.OrderResponsePayload "Order Response Payload"
// @Failure 400 {string} string "invalid request payload"
// @Failure 400 {string} string "invalid product IDs"
// @Failure 404 {string} string "product not exists for ID: <id>"
// @Failure 400 {string} string "product '<name>' (ID: <id>) has only <quantity> quantity available"
// @Failure 500 {string} string "failed to update product quantity"
// @Failure 500 {string} string "failed to process order"
// @Failure 500 {string} string "internal server error"
// @Security BearerAuth
// @Router /order/checkout [post]
func (o *orderController) Checkout(res http.ResponseWriter, req *http.Request) {

	// Get the authenticated user from the request context
	authUser, ok := middlewares.GetAuthUser(req)
	if !ok {
		app.SendErrorResponse(res, http.StatusUnauthorized, "unauthorized access")
		return
	}

	// Validate the request payload
	order, err := validations.Order.Checkout(req)
	if err != nil {
		app.SendErrorResponse(res, http.StatusBadRequest, err.Error())
		return
	}

	// Set the Items products IDS in the productIDs slice
	productIDs := make([]uint, len(order.Items))
	for i, orderItem := range order.Items {
		productIDs[i] = orderItem.ProductID

	}

	// Get products from the database
	products := models.Product.GetProductsByIDs(productIDs)
	if len(products) == 0 {
		app.SendErrorResponse(res, http.StatusBadRequest, "invalid product IDs")
		return
	}

	// Create a map of products with product ID as key for easy lookup
	productsMap := make(map[uint]models.ProductSchema)
	for _, product := range products {
		productsMap[product.ID] = product
	}

	// Order total price
	var orderTotalPrice float64

	// Iterate through order items
	// check product availability and update quantities
	for _, orderItem := range order.Items {
		availableProduct, exists := productsMap[orderItem.ProductID]

		// Check if product does not exist, return an error
		if !exists {
			errMsg := fmt.Sprintf("product not exists for ID: %d", orderItem.ProductID)
			app.SendErrorResponse(res, http.StatusNotFound, errMsg)
			return
		}

		// Check if available product quantity is less than ordered quantity, return an error
		if availableProduct.Quantity < orderItem.Quantity {
			errMsg := fmt.Sprintf("product '%s' (ID: %d) has only %d quantity available", availableProduct.Name, availableProduct.ID, availableProduct.Quantity)
			app.SendErrorResponse(res, http.StatusBadRequest, errMsg)
			return
		}

		// Deduct the ordered quantity from available product quantity, update fails return error
		availableProduct.Quantity -= orderItem.Quantity
		if updatedProduct := models.Product.Update(&availableProduct); updatedProduct.Quantity != availableProduct.Quantity {
			app.SendErrorResponse(res, http.StatusInternalServerError, "failed to update product quantity")
			return
		}

		// Calculate total price for the order item
		orderTotalPrice += availableProduct.Price * float64(orderItem.Quantity)
	}

	// Create the order in the database
	createdOrder := models.Order.Create(&models.OrderSchema{
		UserID:  authUser.ID,
		Status:  "pending",
		Total:   orderTotalPrice,
		Address: authUser.Address,
	})

	if createdOrder.ID == 0 {
		app.SendErrorResponse(res, http.StatusInternalServerError, "failed to process order")
		return
	}

	// Create order items in the database
	for _, orderItem := range order.Items {
		productPrice := productsMap[orderItem.ProductID].Price

		models.OrderItem.Create(&models.OrderItemSchema{
			OrderID:   createdOrder.ID,
			ProductID: orderItem.ProductID,
			Quantity:  orderItem.Quantity,
			Price:     productPrice,
		})
	}

	// Prepare the response payload
	orderResponse := types.OrderResponsePayload{
		OrderID:    createdOrder.ID,
		UserID:     createdOrder.UserID,
		Status:     createdOrder.Status,
		Total:      createdOrder.Total,
		OrderItems: order.Items,
	}

	app.SendSuccessResponse(res, http.StatusCreated, orderResponse)
}

var Order = &orderController{}
