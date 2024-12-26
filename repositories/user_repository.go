package repositories

import (
	"ecommerce-api/config"
	"ecommerce-api/models"
)

func CreateUser(user *models.User) error {
	return config.DB.Create(user).Error
}

func GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := config.DB.Where("email = ?", email).First(&user).Error
	return &user, err
}

func GetUserByID(userID uint) (*models.User, error) {
	var user models.User
	if err := config.DB.First(&user, userID).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
