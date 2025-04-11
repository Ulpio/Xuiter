package services

import (
	"github.com/Ulpio/xuiter.git/internal/config"
	"github.com/Ulpio/xuiter.git/internal/models"
)

func RegisterUser(nome, email, senha string) (uint, error) {
	hashedPass, err := HashPassword(senha)
	if err != nil {
		return 0, err
	}
	user := models.User{
		Username: nome,
		Email:    email,
		Password: hashedPass,
	}

	if err := config.DB.Create(&user).Error; err != nil {
		return 0, err
	}

	return user.ID, nil
}
