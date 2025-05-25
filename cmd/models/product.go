package models

import (
	"time"

	"gorm.io/gorm"
)

// ProductSchema represents a product in the system
// @Description Product details
type ProductSchema struct {
	// ID is the unique identifier for the product
	// @json id
	// @example 1
	ID uint `json:"id" example:"1"`

	// CreatedAt represents the timestamp when the product was added
	// @json createdAt
	// @example "2024-03-19T12:00:00Z"
	CreatedAt time.Time `json:"createdAt" example:"2024-03-19T12:00:00Z"`

	// UpdatedAt represents the last updated timestamp of the product
	// @json updatedAt
	// @example "2024-03-19T12:30:00Z"
	UpdatedAt time.Time `json:"updatedAt" example:"2024-03-19T12:30:00Z"`

	// DeletedAt represents the soft delete timestamp
	// @json deletedAt
	// @example null
	DeletedAt gorm.DeletedAt `json:"deletedAt" swaggerignore:"true"`

	// Name of the product
	// @json name
	// @example "Wireless Headphones"
	Name string `json:"name" example:"Wireless Headphones"`

	// Description provides details about the product
	// @json description
	// @example "High-quality wireless headphones with noise cancellation."
	Description string `json:"description" example:"High-quality wireless headphones with noise cancellation."`

	// Image is the URL to the product image
	// @json image
	// @example "https://example.com/images/headphones.jpg"
	Image string `json:"image" example:"https://example.com/images/headphones.jpg"`

	// Price of the product
	// @json price
	// @example 99.99
	Price float64 `json:"price" example:"99.99"`

	// Quantity available in stock
	// @json quantity
	// @example 50
	Quantity int `json:"quantity" example:"50"`
}

func (p *ProductSchema) Create(product *ProductSchema) *ProductSchema {
	DB.Create(&product)
	return product
}

func (p *ProductSchema) GetByName(name string) *ProductSchema {
	var product ProductSchema
	DB.Where("name = ?", name).First(&product)
	return &product
}

func (p *ProductSchema) GetByID(id uint) *ProductSchema {
	var product ProductSchema
	DB.First(&product, id)
	return &product
}

func (p *ProductSchema) GetAll() []ProductSchema {
	var products []ProductSchema
	DB.Find(&products)
	return products
}

func (p *ProductSchema) GetProductsByIDs(productIDs []uint) []ProductSchema {
	var products []ProductSchema
	DB.Where("id IN ?", productIDs).Find(&products)
	return products
}

func (p *ProductSchema) Update(product *ProductSchema) *ProductSchema {
	DB.Save(&product)
	return product
}

var Product = &ProductSchema{}
