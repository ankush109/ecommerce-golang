package controllers

import (
	"net/http"
	"strconv"

	"github.com/ankush109/ecommerce-go/database"
	"github.com/ankush109/ecommerce-go/models"
	"github.com/gin-gonic/gin"
)

func GetProductById(c *gin.Context) {
	idParams := c.Param("id")
	id, err := strconv.Atoi(idParams)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product Id"})
		return
	}
	var product models.Product
	if err := database.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"product:": product})

}
func GetAllProducts(c *gin.Context) {
	var products []models.Product
	if err := database.DB.Find(&products).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not retrieve products from db"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"products": products})
}

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

	product := models.Product{
		Name:     input.Name,
		Price:    input.Price,
		Quantity: input.Quantity,
		UserID:   userID.(uint),
	}

	if err := database.DB.Create(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product created successfully", "product": product})
}
