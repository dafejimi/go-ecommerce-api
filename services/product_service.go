package services

import (
	"ecommerce-api/models"
	"ecommerce-api/repositories"
)

func CreateProduct(product *models.Product) error {
	return repositories.CreateProduct(product)
}

func GetProducts() ([]models.Product, error) {
	return repositories.GetAllProducts()
}

func GetProductByID(id uint) (*models.Product, error) {
	return repositories.GetProductByID(id)
}

func UpdateProduct(id uint, product *models.Product) error {
	product.ID = id
	return repositories.UpdateProduct(product)
}

func DeleteProduct(id uint) error {
	return repositories.DeleteProduct(id)
}
