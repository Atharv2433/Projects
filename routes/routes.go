package routes

import (
	books "LibraryManagementSystem/handlers/books"
	customer "LibraryManagementSystem/handlers/customer"
	login "LibraryManagementSystem/handlers/login"
	test "LibraryManagementSystem/handlers/test"

	"github.com/gofiber/fiber/v3"
)

func Setup(app *fiber.App) {
	app.Post("/api/login", login.Login)
	app.Get("/api/logout", login.Logout)
	app.Get("/api", test.Greet)
	app.Get("/api/book/:id", books.GetById)
	app.Get("/api/books", books.GetAllBooks)
	app.Get("/api/customer/:id", customer.GetById)
	app.Get("/api/customers", customer.GetAllCustomerS)

}
