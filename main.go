package main

import (
	"os"
	"url-shortner/api/config"
	"url-shortner/api/routes"

	"github.com/gofiber/fiber/v2"
)

func init() {
	// Load environment variables, connect to the database, and perform migrations
	config.LoadENV()
	config.ConnectDB()
	config.AutoMigrate()
}

func main() {
	app := fiber.New()

	// Set up a simple route for testing
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("{active}")
	})

	// Initialize URL routes
	routes.UrlRoutes(app)

	// Retrieve the port from the environment variable or use ":8000" as default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	// Start the server
	if err := app.Listen(":" + port); err != nil {
		// Handle potential errors when starting the server
		panic(err)
	}
}
