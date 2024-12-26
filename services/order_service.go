package services

import (
	"ecommerce-api/models"
	"ecommerce-api/repositories"
)

func CreateOrder(order *models.Order) error {
	return repositories.CreateOrder(order)
}

func GetOrders(userID uint) ([]models.Order, error) {
	return repositories.GetOrdersByUserID(userID)
}

func GetOrderByID(id uint) (*models.Order, error) {
	return repositories.GetOrderByID(id)
}

func UpdateOrderStatus(id uint, order *models.Order) error {
	order.ID = id
	return repositories.UpdateOrderStatus(order)
}

func CancelOrder(id uint) error {
	return repositories.DeleteOrder(id)
}
