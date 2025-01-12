package routes

import (
	"github.com/gofiber/fiber/v2"
	"fiber-app/controllers"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", controllers.Home)
	app.Get("/about", controllers.About)

	app.Get("/user", controllers.GetUser)
	app.Get("/users", controllers.GetUsers)
	app.Get("/user/:id", controllers.GetUserByID)

	app.Get("/search", controllers.Search)
}
