package models

type Test struct {
	ID    uint   `json:"id" gorm:"primaryKey"`
	Name  string `json:"name"`
	Pest  string `json:"pest"`
	Email string `json:"email"`
}