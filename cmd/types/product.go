package types

type ProductPayload struct {
	// Name is the name of the product
	// @json name
	// @example "Wireless Headphones"
	Name        string  `json:"name" example:"Wireless Headphones"`

	// Description provides details about the product
	// @json description
	// @example "High-quality wireless headphones with noise cancellation."
	Description string  `json:"description" example:"High-quality wireless headphones with noise cancellation."`

	// Image is the URL to the product image
	// @json image
	// @example "https://example.com/images/headphones.jpg"
	Image       string  `json:"image" example:"https://example.com/images/headphones.jpg"`

	// Price of the product
	// @json price
	// @example 99.99
	Price       float64 `json:"price" example:"99.99"`

	// Quantity available in stock
	// @json quantity
	// @example 50
	Quantity    int     `json:"quantity" example:"50"`
}