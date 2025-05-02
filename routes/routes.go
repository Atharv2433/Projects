package routes

import (
	books "LibraryManagementSystem/handlers/books"
	customer "LibraryManagementSystem/handlers/customer"
	customerbookinventory "LibraryManagementSystem/handlers/customerbookinventory"
	test "LibraryManagementSystem/handlers/test"

	"github.com/gofiber/fiber/v3"
)

func Setup(app *fiber.App) {
	// app.Post("/api/login", login.Login)
	// app.Get("/api/logout", login.Logout)
	app.Get("/api", test.Greet)
	app.Get("/api/book", books.GetById)
	app.Get("/api/books", books.GetAllBooks)
	app.Get("/api/customer", customer.GetById)
	app.Get("/api/customers", customer.GetAllCustomerS)
	app.Get("/api/history", customerbookinventory.GetprevRecords)
	app.Get("/api/current", customerbookinventory.GetCurrBooks)
	app.Put("/api/inventory", customerbookinventory.UpdateBookStatus)
	app.Post("/api/customer", customer.AddCustomer)

}
