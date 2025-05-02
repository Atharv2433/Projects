package customerbookinventory

import (
	"LibraryManagementSystem/config"
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v3"
)

func UpdateBookStatus(c fiber.Ctx) error {
	if config.Session == nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Database Connection Problem",
		})
	}

	// Define a struct for input
	type UpdateRequest struct {
		Status string `json:"status"`
		Date   string `json:"date"`
		Mode   string `json:"mode"`
		Name   string `json:"name"`
	}

	var req UpdateRequest

	// Manually parse the body
	if err := json.Unmarshal(c.Body(), &req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Execute the SQL query
	query := `UPDATE customer_book_inventory 
	          SET book_status = ? 
	          WHERE mode_of_payment = ? AND customer_name = ? AND purchase_date = ?`

	err := config.Session.Query(query, req.Status, req.Mode, req.Name, req.Date).Exec()
	if err != nil {
		fmt.Println("Error updating status:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update book status",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Book status updated successfully",
	})
}
