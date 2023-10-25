package handlers

import (
	"fmt"
	"golang-db/models"
	"golang-db/services"

	"github.com/gofiber/fiber/v2"
)

type bookHandler struct {
	bookSrv services.BookService
}

func NewBookHandler(bookSrv services.BookService) bookHandler {
	return bookHandler{bookSrv}
}

func (h bookHandler) GetAllBook(c *fiber.Ctx) error {

	res, err := h.bookSrv.GetAllBook()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"massage": err.Error(),
		})
	}

	return c.JSON(res)
}

func (h bookHandler) CreateBook(c *fiber.Ctx) error {

	body := models.HandBookModel{}
	err := c.BodyParser(&body)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"massage": err.Error(),
		})
	}

	err = h.bookSrv.CreateBook(models.SrvBookModel(body))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"massage": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"massage": "create success",
	})
}

func (h bookHandler) UpdateBook(c *fiber.Ctx) error {
	body := models.UpdateHandBookModel{}
	bookID := c.Params("id")
	err := c.BodyParser(&body)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"massage": err.Error(),
		})
	}

	err = h.bookSrv.UpdateBook(bookID, models.UpdateSrvBookModel(body))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"massage": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"massage": "update success",
	})
}

func (h bookHandler) DeleteBook(c *fiber.Ctx) error {
	bookID := c.Params("id")
	err := h.bookSrv.DeleteBook(bookID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"massage": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"massage": "delete success",
	})
}

func validateToken(token string) error {
	if token == "" {
		return fmt.Errorf("token should not be empty")
	}

	return nil
}
