package book

import (
	"fmt"

	"github.com/cokikip/go-fiber-crud/database"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)

type Book struct {
	gorm.Model
	Title  string `json:"title"`
	Auther string `json:"auther"`
	Rating int    `json:"rating"`
}

func GetBooks(c *fiber.Ctx) {
	db := database.DBConn
	var books []Book
	db.Find(&books)
	c.JSON(books)
}

// Get a bool by ID
func GetBook(c *fiber.Ctx) {
	db := database.DBConn
	id := c.Params("id")
	var book Book
	db.Find(&book, id)
	c.JSON(book)
}

// Create a new book from a request body
func NewBook(c *fiber.Ctx) {
	db := database.DBConn
	var book Book
	if err := c.BodyParser(&book); err != nil {
		return
	}
	// book.Auther = "Collins"
	// book.Title = "How to code in go"
	// book.Rating = 5
	db.Create(&book)
	c.JSON(book)
}

// Update a book
func UpdateBook(c *fiber.Ctx) {
	db := database.DBConn
	var book Book
	id := c.Params("id")
	db.First(&book, id)
	if err := c.BodyParser(&book); err != nil {
		return
	}
	db.Update(&book, id)
	c.JSON(book)
}

// Delete a book by ID
func DeleteBook(c *fiber.Ctx) {
	db := database.DBConn
	id := c.Params("id")
	var book Book
	db.First(&book, id)

	if book.Title == "" {
		c.Status(500).Send("No book found with given ID")
		return
	}
	db.Delete(book)
	s := fmt.Sprintf("Book with id %d was successfully deleted", book.ID)
	c.Status(200).Send(s)
}
