package handlers

import "github.com/gofiber/fiber/v3"

func Logout(c fiber.Ctx) error {
	return c.SendString("Logout Page")
}
