// controllers/userController.go

package controllers

import (
	"gorm.io/gorm"
	"fiber-app/models"
)

// PaginationResponse is the struct for the paginated response
type PaginationResponse struct {
	Total    int      `json:"total"`
	Users    []models.User `json:"users"`
	NextPage *int     `json:"next_page"` // Pointer to check for null
}

// getUsers fetches users with pagination from the database
func GetUsers(db *gorm.DB, page, pageSize int) (*PaginationResponse, error) {
	var users []models.User

	// Calculate the total number of users in the database
	var total int64
	if err := db.Model(&models.User{}).Count(&total).Error; err != nil {
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

// CreateUser creates a new user in the database
func CreateUser(db *gorm.DB, user *models.User) error {
	if err := db.Create(user).Error; err != nil {
		return err
	}
	return nil
}
