package models

import "gorm.io/gorm"

type OrderItemSchema struct {
	gorm.Model
	OrderID   uint    `json:"orderId"`
	ProductID uint    `json:"productId"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
}

func (o *OrderItemSchema) Create(orderItem *OrderItemSchema) *OrderItemSchema {
	DBInstance.Create(&orderItem)
	return orderItem
}

var OrderItem = &OrderItemSchema{}
