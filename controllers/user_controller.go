package controllers

import (
	"net/http"

	"ecommerce-api/models"
	"ecommerce-api/services"

	"github.com/gin-gonic/gin"
)

// Register godoc
// @Summary Register a new user
// @Description This endpoint allows users to register by providing their user details
// @Tags Users
// @Accept json
// @Produce json
// @Param user body models.User true "User registration details"
// @Success 200 {object} gin.H{"message": "Registration successful"}
// @Failure 400 {object} gin.H{"error": "Invalid input"}
// @Failure 500 {object} gin.H{"error": "Internal server error"}
// @Router /register [post]
func Register(c *gin.Context) {
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.RegisterUser(&input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Registration successful"})
}

// Login godoc
// @Summary Log in an existing user
// @Description This endpoint allows users to log in by providing their credentials and returns a token
// @Tags Users
// @Accept json
// @Produce json
// @Param login body models.LoginInput true "User login credentials"
// @Success 200 {object} gin.H{"token": "JWT token"}
// @Failure 400 {object} gin.H{"error": "Invalid input"}
// @Failure 401 {object} gin.H{"error": "Unauthorized"}
// @Router /login [post]
func Login(c *gin.Context) {
	var input models.LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := services.LoginUser(&input)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
