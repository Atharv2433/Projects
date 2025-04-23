package handlers

import "github.com/gofiber/fiber/v3"

func Greet(c fiber.Ctx) error {
	return c.SendString("Hello API is Running Good")
}
