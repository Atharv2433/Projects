package customer

import (
	"LibraryManagementSystem/config"
	"LibraryManagementSystem/models"
	"encoding/json"
	"fmt"

	"github.com/gocql/gocql"
	"github.com/gofiber/fiber/v3"
)

func GetById(c fiber.Ctx) error {
	// Check if the database connection exists
	if config.Session == nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Database connection not established",
		})
	}

	// Define a struct to hold the incoming request body
	type CustomerRequest struct {
		Name   string `json:"name"`
		Gender string `json:"gender"`
	}

	var req CustomerRequest

	// Read the body of the request as a byte slice
	body := c.Body()

	// Unmarshal the byte slice into the CustomerRequest struct
	if err := json.Unmarshal(body, &req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Check if essential parameters are provided
	if req.Name == "" || req.Gender == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Both 'name' and 'gender' parameters are required",
		})
	}

	// Log the incoming parameters for debugging
	fmt.Println("Received customer_name:", req.Name)
	fmt.Println("Received customer_gender:", req.Gender)

	var customer models.CustomerInfo

	// Define the query to fetch the customer details from the database
	query := `SELECT customer_name, customer_email, customer_gender, customer_membership_status, customer_mobile, customer_age , delivery_status
	          FROM customer_info WHERE customer_name = ? AND customer_gender = ? LIMIT 1`

	// Execute the query and scan the result into the customer struct
	err := config.Session.Query(query, req.Name, req.Gender).Scan(
		&customer.CustomerName,
		&customer.CustomerEmail,
		&customer.CustomerGender,
		&customer.CustomerMembershipStatus,
		&customer.CustomerMobile,
		&customer.CustomerAge,
		&customer.CustomerDeliveyStatus,
	)

	// Handle errors
	if err != nil {
		fmt.Println("Query error:", err)

		// If no customer is found, return a 'Not Found' response
		if err == gocql.ErrNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Customer not found",
			})
		}

		// For other errors, return a generic error message
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch customer",
		})
	}

	// Debugging: log the fetched customer details
	fmt.Println("Fetched customer:", customer)

	// Return the customer details as a JSON response
	return c.JSON(customer)
}
