package handlers

import (
	"LibraryManagementSystem/config"
	"LibraryManagementSystem/models"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v3"
)

func GetAllBooks(c fiber.Ctx) error {

	if config.Session == nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Database connection not established",
		})
	}

	var books []models.BookInfo

	query := `SELECT * FROM book_info`
	iter := config.Session.Query(query).Iter()

	var book models.BookInfo
	for iter.Scan(
		// &book.BookID,
		&book.BookName,
		&book.BookCategory,
		&book.BookAuthor,
		&book.BookPrice,
		&book.BookLanguage,
		&book.BookPublishedDate,
		&book.BookEdition,
		&book.BookPublisher,
		&book.BookDescription,
	) {
		bookCopy := book
		books = append(books, bookCopy)
	}

	if err := iter.Close(); err != nil {
		log.Println("Iterator error:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch books",
		})
	}

	fmt.Println(books)
	return c.JSON(books)
}
