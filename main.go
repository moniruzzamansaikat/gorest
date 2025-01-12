// main.go

package main

import (
	"log"
	"github.com/gofiber/fiber/v2"
	"fiber-app/routes"
	"fiber-app/models"
	"fiber-app/config"
)

func main() {
	// Set up database connection via the config package
	db := config.SetupDatabase()

	// Automatically migrate the schema
	modelsToMigrate := []interface{}{
		&models.User{},
		&models.Test{},
	}

	for _, model := range modelsToMigrate {
		err := db.AutoMigrate(model)
		if err != nil {
			log.Fatalf("Migration failed for %T: %v", model, err)
		}
	}

	// Set up Fiber app and routes
	app := fiber.New()
	routes.SetupRoutes(app, db)

	// Start the server
	log.Println("Server running on http://localhost:3000")
	if err := app.Listen(":3000"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
