package models

import "gorm.io/gorm"

type BookSchema struct {
	gorm.Model
	Title  string `json:"title"`
	Author string `json:"author"`
	Pages  int    `json:"pages"`
	Price  int    `json:"price"`
}

func (b *BookSchema) Create(book *BookSchema) *BookSchema {
	DBInstance.Create(&book)
	return book
}

func (b *BookSchema) GetByTitle(title string) *BookSchema {
	var book BookSchema
	DBInstance.Where("title = ?", title).First(&book)
	return &book
}

var Book = &BookSchema{}
