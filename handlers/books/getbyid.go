package handlers

import (
	"LibraryManagementSystem/config"
	"LibraryManagementSystem/models"
	"encoding/json"
	"fmt"

	"github.com/gocql/gocql"
	"github.com/gofiber/fiber/v3"
)

func GetById(c fiber.Ctx) error {
	// Check if the database session is not nil (i.e., the connection is established)
	if config.Session == nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Database connection not established",
		})
	}

	// Define a struct to hold the incoming data from the request body
	type BookRequest struct {
		Name     string `json:"name"`
		Category string `json:"category"`
	}

	var req BookRequest

	// Read the body as a byte slice
	body := c.Body()

	// Unmarshal the byte slice into the BookRequest struct
	if err := json.Unmarshal(body, &req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Check if essential parameters are provided
	if req.Name == "" || req.Category == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Both 'name' and 'category' parameters are required",
		})
	}

	// Log the received parameters for debugging
	fmt.Println("Received book_name:", req.Name)
	fmt.Println("Received book_category:", req.Category)

	var book models.BookInfo

	// Define the query to fetch the book from the database based on name and category
	query := `SELECT * FROM book_info WHERE book_name = ? AND book_category = ? LIMIT 1`

	// Execute the query and scan the result into the book struct
	err := config.Session.Query(query, req.Name, req.Category).Scan(
		&book.BookName,
		&book.BookCategory,
		&book.BookAuthor,
		&book.BookPrice,
		&book.BookLanguage,
		&book.BookPublishedDate,
		&book.BookEdition,
		&book.BookPublisher,
		&book.BookDescription,
	)

	// Handle errors
	if err != nil {
		fmt.Println("Error occurred while fetching the book:", err)

		// If no book is found, return a 'Not Found' response
		if err == gocql.ErrNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Book not found",
			})
		}

		// For other errors, return a generic error message
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch book details",
		})
	}

	// Debugging: log the fetched book for verification
	fmt.Println("Fetched book details:", book)

	// Return the book details as a JSON response
	return c.JSON(book)
}
