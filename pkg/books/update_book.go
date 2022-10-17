package books

import (
	"fmt"

	"github.com/Theodule007/go_fiber_api/pkg/common/models"
	"github.com/gofiber/fiber/v2"
)

type UpdateBookRequestBody struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
}

func (h handler) UpdateBook(c *fiber.Ctx) error {
	id := c.Params("id")
	body := UpdateBookRequestBody{}

	// getting request's body
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	fmt.Println("s'ak vin jwenn mwen", body)
	var book models.Book

	if result := h.DB.First(&book, id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}
	fmt.Println("s'ak te nan base de donnees a", book)
	book.Title = body.Title
	book.Author = body.Author
	book.Description = body.Description
	fmt.Println("s'ak al nan base de donnees a")
	// save book
	h.DB.Save(&book)

	return c.Status(fiber.StatusOK).JSON(&book)
}
