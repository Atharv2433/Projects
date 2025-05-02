package customerbookinventory

import (
	"LibraryManagementSystem/config"
	"fmt"

	"github.com/gofiber/fiber/v3"
)

func GetStatus(c fiber.Ctx) error {
	if config.Session == nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Database Connection Not Established",
		})
	}

	name := c.Params("name")
	fmt.Println("Customer Name:", name)

	var status string

	query := `SELECT book_status FROM customer_book_inventory WHERE customer_name = ?`

	err := config.Session.Query(query, name).Consistency(1).Scan(&status)
	if err != nil {
		fmt.Println("Query error:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to fetch previous books",
		})
	}

	fmt.Println("Previous Book IDs:", status)
	return c.JSON(fiber.Map{
		"Book Status": status,
	})
}
