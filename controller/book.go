package controller

import (
	"gobook/models/book"

	"gorm.io/gorm"
)

type Controller struct {
	db *gorm.DB
}

func NewController(db *gorm.DB) *Controller {
	return &Controller{db}
}

func (c *Controller) GetBooks() []book.Book {
	var books []book.Book

	c.db.Find(&books)

	return books
}

func (c *Controller) StoreBook(b map[string]interface{}) book.Book {
	title := b["title"].(string)
	description := b["description"].(string)
	price := b["price"].(int)

	book := book.Book{
		Title:       title,
		Description: description,
		Price:       price,
	}

	c.db.Create(&book)

	return book
}
