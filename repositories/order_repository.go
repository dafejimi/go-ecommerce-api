package repositories

import (
	"ecommerce-api/config"
	"ecommerce-api/models"
)

func CreateOrder(order *models.Order) error {
	return config.DB.Create(order).Error
}

func GetOrdersByUserID(userID uint) ([]models.Order, error) {
	var orders []models.Order
	err := config.DB.Where("user_id = ?", userID).Find(&orders).Error
	return orders, err
}

func GetOrderByID(id uint) (*models.Order, error) {
	var order models.Order
	err := config.DB.First(&order, id).Error
	return &order, err
}

func UpdateOrderStatus(order *models.Order) error {
	return config.DB.Save(order).Error
}

func DeleteOrder(id uint) error {
	return config.DB.Delete(&models.Order{}, id).Error
}
