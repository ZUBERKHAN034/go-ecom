package models

import (
	"time"

	"gorm.io/gorm"
)

// OrderItemSchema represents an item in an order
// @Description Order item details
type OrderItemSchema struct {
	// ID is the unique identifier for the order item
	// @json id
	// @example 1
	ID uint `json:"id" example:"1"`

	// CreatedAt represents the timestamp when the order item was created
	// @json createdAt
	// @example "2024-03-19T12:00:00Z"
	CreatedAt time.Time `json:"createdAt" example:"2024-03-19T12:00:00Z"`

	// UpdatedAt represents the last updated timestamp of the order item
	// @json updatedAt
	// @example "2024-03-19T12:30:00Z"
	UpdatedAt time.Time `json:"updatedAt" example:"2024-03-19T12:30:00Z"`

	// DeletedAt represents the soft delete timestamp
	// @json deletedAt
	// @example null
	DeletedAt gorm.DeletedAt `json:"deletedAt" swaggerignore:"true"`

	// OrderID represents the ID of the associated order
	// @json orderId
	// @example 101
	OrderID uint `json:"orderId" example:"101"`

	// ProductID represents the ID of the purchased product
	// @json productId
	// @example 5001
	ProductID uint `json:"productId" example:"5001"`

	// Quantity of the product in this order item
	// @json quantity
	// @example 2
	Quantity int `json:"quantity" example:"2"`

	// Price of the product in this order item
	// @json price
	// @example 49.99
	Price float64 `json:"price" example:"49.99"`
}

func (o *OrderItemSchema) Create(orderItem *OrderItemSchema) *OrderItemSchema {
	DB.Create(&orderItem)
	return orderItem
}

var OrderItem = &OrderItemSchema{}
