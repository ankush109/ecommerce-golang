package controllers

import (
	"fmt"
	"net/http"

	"github.com/ankush109/ecommerce-go/database"
	"github.com/ankush109/ecommerce-go/models"
	"github.com/gin-gonic/gin"
)

func CreateOrder(c *gin.Context) {
	var orderInput struct {
		ProductID uint `json:"product_id" binding:"required"`
		Quantity  int  `json:"quantity" binding:"required"`
	}
	if err := c.ShouldBindJSON(&orderInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var product models.Product
	if err := database.DB.First(&product, orderInput.ProductID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	totalPrice := float64(orderInput.Quantity) * (product.Price)
	_userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}
	uid := _userID.(uint)
	fmt.Println(uid, "userId...")
	order := models.Order{
		UserID:     uid,
		ProductID:  orderInput.ProductID,
		Quantity:   orderInput.Quantity,
		TotalPrice: totalPrice,
	}
	fmt.Println(order, "order..")
	if err := database.DB.Create(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create Order"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"order": order})

}

func GetOrders(c *gin.Context) {
	var orders []models.Order
	userID, _ := c.Get("userID")

	if err := database.DB.Where("user_id = ?", userID).Preload("Product").Find(&orders).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve orders"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"orders": orders})
}
