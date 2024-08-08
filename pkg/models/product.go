package models

import "gorm.io/gorm"

type ProductSchema struct {
	gorm.Model
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Image       string  `json:"image"`
	Price       float64 `json:"price"`
	Quantity    int     `json:"quantity"`
}

func (p *ProductSchema) Create(product *ProductSchema) *ProductSchema {
	DBInstance.Create(&product)
	return product
}

func (p *ProductSchema) GetByName(name string) *ProductSchema {
	var product ProductSchema
	DBInstance.Where("name = ?", name).First(&product)
	return &product
}

func (p *ProductSchema) GetByID(id uint) *ProductSchema {
	var product ProductSchema
	DBInstance.First(&product, id)
	return &product
}

func (p *ProductSchema) GetAll() []ProductSchema {
	var products []ProductSchema
	DBInstance.Find(&products)
	return products
}

var Product = &ProductSchema{}
