package main

import (
	"github.com/ankush109/ecommerce-go/database"
	"github.com/ankush109/ecommerce-go/models"
	"github.com/ankush109/ecommerce-go/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	database.ConnectDatabase()

	database.DB.AutoMigrate(&models.User{})
	database.DB.AutoMigrate(&models.Product{})
	database.DB.AutoMigrate(&models.Order{})

	database.ConnectDatabase()

	routes.SetupRoutes(r)

	r.Run(":8080")
}
