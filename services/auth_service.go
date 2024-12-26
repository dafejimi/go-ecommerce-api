package services

import (
	"ecommerce-api/models"
	"ecommerce-api/repositories"
	"ecommerce-api/utils"
	"errors"
)

func RegisterUser(user *models.User) error {
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = hashedPassword
	return repositories.CreateUser(user)
}

func LoginUser(input *models.LoginInput) (string, error) {
	user, err := repositories.GetUserByEmail(input.Email)
	if err != nil {
		return "", errors.New("user not found")
	}

	if !utils.CheckPassword(input.Password, user.Password) {
		return "", errors.New("incorrect password")
	}

	return utils.GenerateJWT(user.ID, user.IsAdmin)
}

func IsAdmin(userID uint) bool {
	// Query the user from the database by userID
	user, err := repositories.GetUserByID(userID) // Assuming this service function exists
	if err != nil {
		return false // Return false if user not found or there's an error
	}

	// Check if the user is an admin using the is_admin field
	return user.IsAdmin
}
