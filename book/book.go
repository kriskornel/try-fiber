package book

import "github.com/gofiber/fiber/v2"

func GetBooks(c *fiber.Ctx) {
	c.Send("All books")
}

func GetBook(c *fiber.Ctx) {
	c.Send("Single Book")
}

func NewBook(c *fiber.Ctx) {
	c.Send("New Book")
}

func DeleteBook(c *fiber.Ctx) {
	c.Send("Delete Book")
}
