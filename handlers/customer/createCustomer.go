package customer

import (
	"LibraryManagementSystem/config"
	"LibraryManagementSystem/models"
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v3"
)

func AddCustomer(c fiber.Ctx) error {
	// Check if the database connection exists
	if config.Session == nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Database connection not established",
		})
	}

	// Define a struct to bind the incoming request body to
	var req models.CustomerInfo

	// Read the body as a byte slice
	body := c.Body()

	// Unmarshal the byte slice into the CustomerInfo struct
	if err := json.Unmarshal(body, &req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Validate required fields
	if req.CustomerName == "" || req.CustomerEmail == "" || req.CustomerMobile == "" || req.CustomerGender == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Required fields missing (name, email, mobile, or gender)",
		})
	}

	// Prepare the query to insert the customer into the table
	query := `INSERT INTO library2.customer_info (customer_gender, customer_name, customer_age, customer_email, 
	          customer_membership_status, customer_mobile, delivery_status) 
	          VALUES (?, ?, ?, ?, ?, ?, ?)`

	// Execute the query to insert the customer
	err := config.Session.Query(query, req.CustomerGender, req.CustomerName, req.CustomerAge, req.CustomerEmail,
		req.CustomerMembershipStatus, req.CustomerMobile, req.CustomerDeliveyStatus).Exec()

	if err != nil {
		fmt.Println("Error while adding customer:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to add customer",
		})
	}

	// Return a success message
	return c.JSON(fiber.Map{
		"message":  "Customer added successfully",
		"customer": req,
	})
}
