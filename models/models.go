// models/user.go

package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"unique" json:"username"`
	Email    string `gorm:"unique" json:"email"`
	Password string `json:"-"`
}

type Product struct {
	gorm.Model
	ID       uint    `gorm:"primaryKey" json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
	UserID   uint    `json:"user_id"` // Foreign key to associate with the user
}
