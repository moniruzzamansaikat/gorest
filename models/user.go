package models

type User struct {
	ID           uint   `json:"id" gorm:"primaryKey"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	Email        string `json:"email"`
	Test         string `json:"test"`
}