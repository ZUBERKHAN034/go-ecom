package types

type OrderItemPayload struct {
	// ProductID is the unique identifier for the product
	// @json productId
	// @example 1
	ProductID uint `json:"productId" example:"1"`

	// Quantity is the number of units of the product in the order
	// @json quantity
	// @example 2
	Quantity int `json:"quantity" example:"2"`
}

type OrderPayload struct {
	// List of order items
	// @json orderItems
	Items []OrderItemPayload `json:"orderItems"`
}

type OrderResponsePayload struct {
	// OrderID is the unique identifier for the order
	// @json orderId
	// @example 12345
	OrderID uint `json:"orderId" example:"12345"`

	// UserID is the unique identifier for the user who placed the order
	// @json userId
	// @example 67890
	UserID uint `json:"userId" example:"67890"`

	// Status is the current status of the order
	// @json status
	// @example "pending"
	Status string `json:"status" example:"pending"`

	// Total is the total cost of the order
	// @json total
	// @example 99.99
	Total float64 `json:"total" example:"99.99"`

	// order Items is the list of order items in the order
	// @json orderItems
	OrderItems []OrderItemPayload `json:"orderItems"`
}
