package customerbookinventory

import (
	"LibraryManagementSystem/config"
	"fmt"

	"github.com/gocql/gocql"
	"github.com/gofiber/fiber/v3"
)

func GetCurrBooks(c fiber.Ctx) error {
	if config.Session == nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Database Connection Problem",
		})
	}

	// Get book status and customer name from parameters
	status := c.Params("status")
	customerName := c.Params("name")

	// Prepare the query to fetch the current books
	query := `SELECT current_books FROM customer_book_inventory WHERE book_status = ? AND customer_name = ?`

	var listCurrentBooks []string

	err := config.Session.Query(query, status, customerName).Consistency(gocql.One).Scan(&listCurrentBooks)
	if err != nil {
		if err == gocql.ErrNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Customer or books not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Query execution failed",
		})
	}

	// Return the current books as a list of book IDs
	fmt.Println(listCurrentBooks)

	return c.JSON(fiber.Map{
		"list_current_books": listCurrentBooks,
	})
}
