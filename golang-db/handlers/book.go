package handlers

import "github.com/gofiber/fiber/v2"

type BookHandler interface {
	GetAllBook(c *fiber.Ctx) error
	CreateBook(c *fiber.Ctx) error
	UpdateBook(c *fiber.Ctx) error
	DeleteBook(c *fiber.Ctx) error
}
