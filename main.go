package main

import (
	"log"
	"github.com/gofiber/fiber/v2"
	"fiber-app/routes"
)

func main() {
	app := fiber.New()

	routes.SetupRoutes(app)

	log.Println("Server running on http://localhost:3000")
	if err := app.Listen(":3000"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
