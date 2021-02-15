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
