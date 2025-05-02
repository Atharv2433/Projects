package customerbookinventory

import (
	"LibraryManagementSystem/config"
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v3"
)

func GetprevRecords(c fiber.Ctx) error {
	if config.Session == nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Database Connection Not Established",
		})
	}

	// Define struct for request body
	type RequestBody struct {
		Name   string `json:"name"`
		Status string `json:"status"`
	}

	var req RequestBody

	// Unmarshal the request body manually
	if err := json.Unmarshal(c.Body(), &req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	fmt.Println("Status:", req.Status)
	fmt.Println("Customer Name:", req.Name)

	var previousBooks []string

	query := `SELECT previous_books FROM customer_book_inventory WHERE book_status = ? AND customer_name = ?`

	iter := config.Session.Query(query, req.Status, req.Name).Iter()
	var book string

	for iter.Scan(&book) {
		previousBooks = append(previousBooks, book)
	}

	if err := iter.Close(); err != nil {
		fmt.Println("Query error:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to fetch previous books",
		})
	}

	return c.JSON(fiber.Map{
		"book_ids": previousBooks,
	})
}
