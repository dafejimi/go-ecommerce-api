package main

import (
	"ecommerce-api/config"
	"ecommerce-api/controllers"
	"ecommerce-api/middleware"

	_ "ecommerce-api/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

// @title E-commerce API
// @version 1.0
// @description RESTful API for e-commerce platform

// @contact.email jimijay.oj@gmailcom

// @host localhost:8080
// @BasePath /api
func main() {
	config.LoadEnv()
	r := gin.Default()

	// Middleware
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Routes
	api := r.Group("/api")
	{
		api.POST("/register", controllers.Register)
		api.POST("/login", controllers.Login)

		products := api.Group("/products")
		products.Use(middleware.AuthMiddleware(), middleware.AdminMiddleware())
		{
			products.POST("", controllers.CreateProduct)
			products.GET("", controllers.GetProducts)
			products.GET("/:id", controllers.GetProductByID)
			products.PUT("/:id", controllers.UpdateProduct)
			products.DELETE("/:id", controllers.DeleteProduct)
		}

		orders := api.Group("/orders")
		orders.Use(middleware.AuthMiddleware())
		{
			orders.POST("", controllers.CreateOrder)
			orders.GET("", controllers.GetOrders)
			orders.GET("/:id", controllers.GetOrderByID)
			orders.PUT("/:id", controllers.UpdateOrderStatus)
			orders.DELETE("/:id", controllers.CancelOrder)
		}
	}

	// Run the server
	r.Run()
}
