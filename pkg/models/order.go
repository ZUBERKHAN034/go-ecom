package models

import "gorm.io/gorm"

type OrderSchema struct {
	gorm.Model
	UserID   uint    `json:"userId"`
	Quantity int     `json:"quantity"`
	Status   int     `json:"status"`
	Address  string  `json:"address"`
	Total    float64 `json:"total"`
}

func (o *OrderSchema) Create(order *OrderSchema) *OrderSchema {
	DBInstance.Create(&order)
	return order
}

var Cart = &OrderSchema{}
