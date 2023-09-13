package service

import (
	"errors"
	"fiber_postgres/database"
	"fiber_postgres/handler"
	"fiber_postgres/model"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetAllBooks(c *fiber.Ctx) error {
	db := database.DB
	var books []model.Book
	db.Find(&books)
	return c.JSON(books)
}

func GetBook(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB
	var book model.Book
	err := db.First(&book, id).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return handler.EntityNotFound("No book found")
	} else if err != nil {
		return handler.Exception(err.Error())
	}
	return c.JSON(book)
}

func CreateBook(c *fiber.Ctx) error {
	db := database.DB
	book := new(model.Book)
	if err := c.BodyParser(book); err != nil {
		return handler.BadRequest("Invalid param")
	}
	db.Create(book)
	return c.JSON(book)
}

func UpdateBook(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB
	var book model.Book
	err := db.First(&book, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return handler.EntityNotFound("No book found")
	} else if err != nil {
		return handler.Exception(err.Error())
	}
	updatedBook := new(model.Book)

	if err := c.BodyParser(updatedBook); err != nil {
		return handler.BadRequest("Invalid params")
	}

	updatedBook = &model.Book{Title: updatedBook.Title, Author: updatedBook.Author, Rating: updatedBook.Rating}

	if err = db.Model(&book).Updates(updatedBook).Error; err != nil {
		return handler.Exception(err.Error())
	}
	return c.SendStatus(204)

}

func DeleteBook(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB
	var book model.Book
	err := db.First(&book, id).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return handler.EntityNotFound("No book found")
	} else if err != nil {
		return handler.Exception(err.Error())
	}
	db.Delete(&book)
	return c.SendStatus(204)
}
