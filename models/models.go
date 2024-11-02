package models

import (
	"time"

	"gorm.io/gorm"
)

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
	UserID   uint    `json:"user_id"`
}

type Order struct {
	gorm.Model
	ID         uint      `gorm:"primaryKey"`
	UserID     uint      `json:"user_id"`              // Foreign key to the User model
	ProductID  uint      `json:"product_id"`           // Foreign key ID for Product
	Product    Product   `gorm:"foreignKey:ProductID"` // Automatically populated based on ProductID
	Quantity   int       `json:"quantity"`
	TotalPrice float64   `json:"total_price"`
	CreatedAt  time.Time `json:"created_at"`
}
