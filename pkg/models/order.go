package models

import (
	"time"

	"gorm.io/gorm"
)

// OrderSchema represents an order in the system
// @Description Order details
type OrderSchema struct {
	// ID is the unique identifier for the order
	// @json id
	// @example 1
	ID uint `json:"id" example:"1"`

	// CreatedAt represents the order creation timestamp
	// @json createdAt
	// @example "2024-03-19T12:00:00Z"
	CreatedAt time.Time `json:"createdAt" example:"2024-03-19T12:00:00Z"`

	// UpdatedAt represents the last update timestamp of the order
	// @json updatedAt
	// @example "2024-03-19T12:30:00Z"
	UpdatedAt time.Time `json:"updatedAt" example:"2024-03-19T12:30:00Z"`

	// DeletedAt represents the soft delete timestamp
	// @json deletedAt
	// @example null
	DeletedAt gorm.DeletedAt `json:"deletedAt" swaggerignore:"true"`

	// UserID represents the ID of the user who placed the order
	// @json userId
	// @example 123
	UserID uint `json:"userId" example:"123"`

	// Quantity of the items in the order
	// @json quantity
	// @example 2
	Quantity int `json:"quantity" example:"2"`

	// Status of the order (e.g., 0 = Pending, 1 = Shipped, 2 = Delivered)
	// @json status
	// @example 1
	Status int `json:"status" example:"1"`

	// Address where the order needs to be delivered
	// @json address
	// @example "123, Street Name, City, Country"
	Address string `json:"address" example:"123, Street Name, City, Country"`

	// Total cost of the order
	// @json total
	// @example 99.99
	Total float64 `json:"total" example:"99.99"`
}

func (o *OrderSchema) Create(order *OrderSchema) *OrderSchema {
	DBInstance.Create(&order)
	return order
}

var Order = &OrderSchema{}
