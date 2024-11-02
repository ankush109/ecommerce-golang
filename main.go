// main.go

package main

import (
	"github.com/ankush109/ecommerce-go/database"
	"github.com/ankush109/ecommerce-go/models"
	"github.com/ankush109/ecommerce-go/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Connect to the database
	database.ConnectDatabase()

	// Run migrations to create tables based on models
	database.DB.AutoMigrate(&models.User{})
	database.DB.AutoMigrate(&models.Product{})
	database.DB.AutoMigrate(&models.Order{})

	// Connect to the database
	database.ConnectDatabase()

	// Set up routes
	routes.SetupRoutes(r)

	// Start the server
	r.Run(":8080")
}
