package routes

import (
	"log"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// User model for GORM
type User struct {
	ID    uint   `json:"id" gorm:"primaryKey"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// PaginationResponse is the struct for the paginated response
type PaginationResponse struct {
	Total    int    `json:"total"`
	Users    []User `json:"users"`
	NextPage *int   `json:"next_page"` // Pointer to check for null
}

// Function to fetch users with pagination from the database using GORM
func getUsers(db *gorm.DB, page, pageSize int) (*PaginationResponse, error) {
	var users []User

	// Calculate the total number of users in the database
	var total int64
	if err := db.Model(&User{}).Count(&total).Error; err != nil {
		return nil, err
	}

	// Fetch the users with the specified page and page size
	if err := db.Offset((page - 1) * pageSize).Limit(pageSize).Find(&users).Error; err != nil {
		return nil, err
	}

	// Check if there's a next page
	var nextPage *int
	if page*pageSize < int(total) {
		nextPage = &[]int{page + 1}[0]
	}

	// Return the paginated response
	return &PaginationResponse{
		Total:    int(total),
		Users:    users,
		NextPage: nextPage,
	}, nil
}

// Function to set up routes
func SetupRoutes(app *fiber.App, db *gorm.DB) {
	// Basic GET route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Fiber with GORM!")
	})

	// Route that returns a list of users with pagination from MySQL database using GORM
	app.Get("/users", func(c *fiber.Ctx) error {
		// Get pagination parameters from query string
		page := c.QueryInt("page", 1)     // Default to page 1
		pageSize := c.QueryInt("page_size", 10) // Default to 10 users per page

		paginatedUsers, err := getUsers(db, page, pageSize)
		if err != nil {
			log.Println("Failed to fetch users from database:", err)
			return c.Status(500).SendString("Failed to fetch users from database")
		}

		return c.JSON(paginatedUsers)
	})

	// Create a new user (for testing purposes)
	app.Post("/user", func(c *fiber.Ctx) error {
		var newUser User
		if err := c.BodyParser(&newUser); err != nil {
			return c.Status(400).SendString("Invalid request body")
		}

		// Save the user to the database
		if err := db.Create(&newUser).Error; err != nil {
			return c.Status(500).SendString("Failed to create user")
		}

		return c.JSON(newUser)
	})
}
