package routes

import (
	"github.com/ankush109/ecommerce-go/controllers"
	"github.com/ankush109/ecommerce-go/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {

	router.GET("/ping", controllers.HealthCheck)
	router.POST("/register", controllers.Register)
	router.POST("/login", controllers.Login)
	router.GET("/products/:id", controllers.GetProductById)
	productGroup := router.Group("/products")
	productGroup.Use(middlewares.JWTMiddleware())
	{
		productGroup.POST("/", controllers.CreateProduct)

	}

	orderGroup := router.Group("/orders")
	orderGroup.Use(middlewares.JWTMiddleware())
	{
		orderGroup.POST("/", controllers.CreateOrder)
		orderGroup.GET("/", controllers.GetOrders)
	}
}
