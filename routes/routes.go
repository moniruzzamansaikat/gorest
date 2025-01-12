// routes/routes.go

package routes

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"fiber-app/models"
	"fiber-app/controllers"
)

// SetupRoutes configures the routes for the application
func SetupRoutes(app *fiber.App, db *gorm.DB) {
	// Basic route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Fiber with GORM!")
	})

	// Route to fetch users with pagination
	app.Get("/users", func(c *fiber.Ctx) error {
		page := c.QueryInt("page", 1)     // Default to page 1
		pageSize := c.QueryInt("page_size", 10) // Default to 10 users per page

		paginatedUsers, err := controllers.GetUsers(db, page, pageSize)
		if err != nil {
			return c.Status(500).SendString("Failed to fetch users from database")
		}

		return c.JSON(paginatedUsers)
	})

	// Route to create a new user
	app.Post("/user", func(c *fiber.Ctx) error {
		var newUser models.User
		if err := c.BodyParser(&newUser); err != nil {
			return c.Status(400).SendString("Invalid request body")
		}

		if err := controllers.CreateUser(db, &newUser); err != nil {
			return c.Status(500).SendString("Failed to create user")
		}

		return c.JSON(newUser)
	})
}
