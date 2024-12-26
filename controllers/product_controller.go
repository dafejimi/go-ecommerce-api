package controllers

import (
	"net/http"
	"strconv"

	"ecommerce-api/models"
	"ecommerce-api/services"

	"github.com/gin-gonic/gin"
)

// CreateProduct godoc
// @Summary Create a new product
// @Description This endpoint allows an admin to create a new product by providing product details
// @Tags Products
// @Accept json
// @Produce json
// @Param product body models.Product true "Product details"
// @Success 200 {object} gin.H{"message": "Product created successfully"}
// @Failure 400 {object} gin.H{"error": "Invalid input"}
// @Failure 500 {object} gin.H{"error": "Internal server error"}
// @Router /products [post]
func CreateProduct(c *gin.Context) {
	var input models.Product
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.CreateProduct(&input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product created successfully"})
}

// GetProducts godoc
// @Summary Get all products
// @Description This endpoint allows anyone to fetch a list of all available products
// @Tags Products
// @Accept json
// @Produce json
// @Success 200 {object} gin.H{"products": []models.Product}
// @Failure 500 {object} gin.H{"error": "Internal server error"}
// @Router /products [get]
func GetProducts(c *gin.Context) {
	products, err := services.GetProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"products": products})
}

// GetProductByID godoc
// @Summary Get a product by ID
// @Description This endpoint allows anyone to fetch details of a specific product by its ID
// @Tags Products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} gin.H{"product": models.Product}
// @Failure 400 {object} gin.H{"error": "Invalid product ID"}
// @Failure 404 {object} gin.H{"error": "Product not found"}
// @Failure 500 {object} gin.H{"error": "Internal server error"}
// @Router /products/{id} [get]
func GetProductByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	product, err := services.GetProductByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"product": product})
}

// UpdateProduct godoc
// @Summary Update an existing product
// @Description This endpoint allows an admin to update the details of an existing product
// @Tags Products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Param product body models.Product true "Updated product details"
// @Success 200 {object} gin.H{"message": "Product updated successfully"}
// @Failure 400 {object} gin.H{"error": "Invalid input"}
// @Failure 404 {object} gin.H{"error": "Product not found"}
// @Failure 500 {object} gin.H{"error": "Internal server error"}
// @Router /products/{id} [put]
func UpdateProduct(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var input models.Product
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.UpdateProduct(uint(id), &input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product updated successfully"})
}

// DeleteProduct godoc
// @Summary Delete a product
// @Description This endpoint allows an admin to delete an existing product by its ID
// @Tags Products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} gin.H{"message": "Product deleted successfully"}
// @Failure 404 {object} gin.H{"error": "Product not found"}
// @Failure 500 {object} gin.H{"error": "Internal server error"}
// @Router /products/{id} [delete]
func DeleteProduct(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	if err := services.DeleteProduct(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}
