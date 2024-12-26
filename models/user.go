package models

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	IsAdmin  bool   `gorm:"default:false"`
}

type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}
