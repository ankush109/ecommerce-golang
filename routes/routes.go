// routes/routes.go

package routes

import (
	"github.com/ankush109/ecommerce-go/controllers"
	"github.com/gin-gonic/gin"
)

// SetupRoutes sets up the routes for the application
func SetupRoutes(router *gin.Engine) {
	// Health check route
	router.GET("/ping", controllers.HealthCheck)
	router.POST("/register", controllers.Register)
	//r.POST("/login", controllers.Login)

	// You can add more routes here, for example:
	// router.GET("/users", controllers.GetUsers)
	// router.POST("/users", controllers.CreateUser)
}
