package book

import "github.com/gofiber/fiber/v2"

func GetBooks(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"data": "All books",
	})
}

func GetBook(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Ok!",
	})
}

func NewBook(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Ok!",
	})
}

func DeleteBook(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Ok!",
	})
}
