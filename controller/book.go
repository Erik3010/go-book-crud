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

func (c Controller) GetBooks() []book.Book {
	var books []book.Book

	c.db.Find(&books)

	return books
}

func (c Controller) StoreBook(bookData *book.Book) *gorm.DB {
	result := c.db.Create(&bookData)

	return result
}

func (c Controller) ShowBook(id int) (bookData book.Book, err error) {
	result := c.db.First(&bookData, id)

	return bookData, result.Error
}

func (c Controller) UpdateBook(bookData *book.Book) *gorm.DB {
	result := c.db.Model(&bookData).Select("*").Omit("created_at").Updates(bookData)
	return result
}

func (c Controller) DeleteBook(id int) *gorm.DB {
	result := c.db.Delete(&book.Book{}, id)
	return result
}
