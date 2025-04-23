package handlers

import (
	"LibraryManagementSystem/config"
	"LibraryManagementSystem/models"

	"github.com/gocql/gocql"
	"github.com/gofiber/fiber/v3"
)

func GetById(c fiber.Ctx) error {
	// Check if the session is initialized
	if config.Session == nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Database connection not established",
		})
	}

	id := c.Params("id")
	bookId, err := gocql.ParseUUID(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid UUID format",
		})
	}

	var book models.BookInfo
	query := `SELECT * FROM book_info WHERE book_id = ?`
	err = config.Session.Query(query, bookId).Scan(
		&book.BookID,
		&book.BookName,
		&book.Genre,
		&book.BookAuthor,
		&book.BookPrice,
		&book.BookLanguage,
		&book.BookPublishedDate,
		&book.BookEdition,
		&book.BookPublisher,
		&book.BookDescription,
	)

	if err != nil {
		if err == gocql.ErrNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Book not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch book",
		})
	}

	return c.JSON(book)
}
