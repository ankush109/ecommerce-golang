// controllers/product.go

package controllers

import (
	"net/http"

	"github.com/ankush109/ecommerce-go/database"
	"github.com/ankush109/ecommerce-go/models"
	"github.com/gin-gonic/gin"
)

// CreateProduct handles the creation of a new product
func CreateProduct(c *gin.Context) {
	var input models.Product
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	// Create the product and associate it with the user
	product := models.Product{
		Name:     input.Name,
		Price:    input.Price,
		Quantity: input.Quantity,
		UserID:   userID.(uint), // Cast userID to uint
	}

	// Save the product to the database
	if err := database.DB.Create(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product created successfully", "product": product})
}
