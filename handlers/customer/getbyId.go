package customer

import (
	"LibraryManagementSystem/config"
	"LibraryManagementSystem/models"

	"github.com/gocql/gocql"
	"github.com/gofiber/fiber/v3"
)

func GetById(c fiber.Ctx) error {
	// return c.SendString("Customer By Id")

	if config.Session == nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Database Connection not Done",
		})
	}
	id := c.Params("id")
	customerId, err := gocql.ParseUUID(id)
	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"error": "Invalid Customer Id",
		})
	}

	var customer models.CustomerInfo

	query := `SELECT * FROM customer_info where customer_id = ?`
	err = config.Session.Query(query, customerId).Scan(
		&customer.CustomerID,
		&customer.CustomerAddr,
		&customer.CustomerEmail,
		&customer.CustomerGender,
		&customer.CustomerGSTNumber,
		&customer.CustomerMembershipStatus,
		&customer.CustomerMobile,
		&customer.CustomerName,
		&customer.CustomerPassword,
		&customer.CustomerUsername,
		&customer.CustomerPreferredLanguage,
	)

	if err != nil {
		if err == gocql.ErrNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Customer not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch Customer",
		})
	}

	return c.JSON(customer)

}
