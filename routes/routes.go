// routes/routes.go

package routes

import (
	"github.com/ankush109/ecommerce-go/controllers"
	"github.com/ankush109/ecommerce-go/middlewares"
	"github.com/gin-gonic/gin"
)

// SetupRoutes sets up the routes for the application
func SetupRoutes(router *gin.Engine) {
	// Public routes
	router.GET("/ping", controllers.HealthCheck)
	router.POST("/register", controllers.Register)
	router.POST("/login", controllers.Login)

	// Protected routes
	productGroup := router.Group("/products")
	productGroup.Use(middlewares.JWTMiddleware()) // Apply JWT middleware only to this group
	{
		productGroup.POST("/", controllers.CreateProduct)
		// Additional protected product routes can go here
	}

	// Add more protected groups if needed
	// For example, an orders group:
	// ordersGroup := router.Group("/orders")
	// ordersGroup.Use(middlewares.JWTMiddleware())
	// ordersGroup.POST("/", controllers.CreateOrder)
}
