package controllers

import "github.com/gofiber/fiber/v2"

func GetUser(c *fiber.Ctx) error {
	user := map[string]interface{}{
		"id":    1,
		"name":  "John Doe",
		"email": "john.doe@example.com",
	}
	return c.JSON(user)
}

func GetUsers(c *fiber.Ctx) error {
	users := []map[string]interface{}{
		{"id": 1, "name": "John Doe", "email": "john.doe@example.com"},
		{"id": 2, "name": "Jane Smith", "email": "jane.smith@example.com"},
		{"id": 3, "name": "Alice Johnson", "email": "alice.johnson@example.com"},
	}
	return c.JSON(users)
}

func GetUserByID(c *fiber.Ctx) error {
	id := c.Params("id")
	return c.SendString("User ID: " + id)
}
