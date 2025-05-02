package customer

import (
	"LibraryManagementSystem/config"
	"LibraryManagementSystem/models"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v3"
)

func GetAllCustomerS(c fiber.Ctx) error {
	if config.Session == nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Database Connection Not Established",
		})
	}

	var customers []models.CustomerInfo
	query := `SELECT customer_name, customer_membership_status, customer_age, customer_mobile, customer_email, customer_gender , delivery_status  FROM customer_info`
	iter := config.Session.Query(query).Iter()

	var customer models.CustomerInfo
	for iter.Scan(
		&customer.CustomerName,
		&customer.CustomerMembershipStatus,
		&customer.CustomerAge,
		&customer.CustomerMobile,
		&customer.CustomerEmail,
		&customer.CustomerGender,
		&customer.CustomerDeliveyStatus,
	) {
		copyCustomer := customer
		customers = append(customers, copyCustomer)
	}

	if err := iter.Close(); err != nil {
		log.Println("Iterator Error:", err)
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"error": "Unable to fetch customers",
		})
	}

	fmt.Println(customers)
	return c.JSON(customers)
}
