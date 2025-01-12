package controllers

import "github.com/gofiber/fiber/v2"

// Search handles search query parameters
func Search(c *fiber.Ctx) error {
	query := c.Query("q", "default") // Default to "default" if no query parameter is passed
	return c.SendString("Search query: " + query)
}
