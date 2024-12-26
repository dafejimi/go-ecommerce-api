package controllers

import (
	"net/http"
	"strconv"

	"ecommerce-api/models"
	"ecommerce-api/services"

	"github.com/gin-gonic/gin"
)

// CreateOrder godoc
// @Summary Create a new order
// @Description This endpoint allows a user to place a new order by providing order details
// @Tags Orders
// @Accept json
// @Produce json
// @Param order body models.Order true "Order details"
// @Success 200 {object} gin.H{"message": "Order placed successfully"}
// @Failure 400 {object} gin.H{"error": "Invalid input"}
// @Failure 500 {object} gin.H{"error": "Internal server error"}
// @Router /orders [post]
func CreateOrder(c *gin.Context) {
	var input models.Order
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.CreateOrder(&input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order placed successfully"})
}

// GetOrders godoc
// @Summary Get all orders for the authenticated user
// @Description This endpoint allows a user to retrieve all their orders
// @Tags Orders
// @Accept json
// @Produce json
// @Success 200 {object} gin.H{"orders": []models.Order}
// @Failure 500 {object} gin.H{"error": "Internal server error"}
// @Router /orders [get]
func GetOrders(c *gin.Context) {
	userID := c.MustGet("userID").(uint)
	orders, err := services.GetOrders(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"orders": orders})
}

// GetOrderByID godoc
// @Summary Get a specific order by ID
// @Description This endpoint allows a user to fetch a specific order by its ID
// @Tags Orders
// @Accept json
// @Produce json
// @Param id path int true "Order ID"
// @Success 200 {object} gin.H{"order": models.Order}
// @Failure 400 {object} gin.H{"error": "Invalid order ID"}
// @Failure 404 {object} gin.H{"error": "Order not found"}
// @Failure 500 {object} gin.H{"error": "Internal server error"}
// @Router /orders/{id} [get]
func GetOrderByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	order, err := services.GetOrderByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"order": order})
}

// UpdateOrderStatus godoc
// @Summary Update the status of an existing order
// @Description This endpoint allows an admin or user to update the status of an order
// @Tags Orders
// @Accept json
// @Produce json
// @Param id path int true "Order ID"
// @Param order body models.Order true "Updated order status"
// @Success 200 {object} gin.H{"message": "Order status updated successfully"}
// @Failure 400 {object} gin.H{"error": "Invalid input"}
// @Failure 404 {object} gin.H{"error": "Order not found"}
// @Failure 500 {object} gin.H{"error": "Internal server error"}
// @Router /orders/{id} [put]
func UpdateOrderStatus(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var input models.Order
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.UpdateOrderStatus(uint(id), &input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order status updated successfully"})
}

// CancelOrder godoc
// @Summary Cancel an existing order
// @Description This endpoint allows a user or admin to cancel an order by its ID
// @Tags Orders
// @Accept json
// @Produce json
// @Param id path int true "Order ID"
// @Success 200 {object} gin.H{"message": "Order cancelled successfully"}
// @Failure 404 {object} gin.H{"error": "Order not found"}
// @Failure 500 {object} gin.H{"error": "Internal server error"}
// @Router /orders/{id} [delete]
func CancelOrder(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	if err := services.CancelOrder(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order cancelled successfully"})
}
