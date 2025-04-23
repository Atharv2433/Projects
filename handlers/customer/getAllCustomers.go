package customer

import (
	"LibraryManagementSystem/config"
	"LibraryManagementSystem/models"
	"log"

	"github.com/gofiber/fiber/v3"
)

func GetAllCustomerS(c fiber.Ctx) error {
	// return c.SendString("ALL Customers")

	if config.Session == nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Databse Connection Not Established",
		})
	}

	var customers []models.CustomerInfo

	query := `SELECT * FROM customer_info`
	iter := config.Session.Query(query).Iter()

	var customer models.CustomerInfo
	for iter.Scan(
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
	) {
		copyCustomer := customer
		customers = append(customers, copyCustomer)
	}

	if err := iter.Close(); err != nil {
		log.Println("Iterator Error ....", err)

		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"error": "unable to Fetch",
		})
	}
	return c.JSON(customers)
}
