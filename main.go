package main

import (
	"LibraryManagementSystem/config"
	"LibraryManagementSystem/routes"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Print message indicating the app is starting
	fmt.Println("Starting Library Management ......")

	// Initialize the Fiber app
	app := fiber.New()

	// Establish DB connection
	config.ConnectionDB()

	// Set up routes
	routes.Setup(app)

	// Start the server
	log.Fatal(app.Listen(":3000"))
}
