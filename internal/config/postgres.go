package config

import (
	"fmt"
	"os"

	"github.com/Ulpio/xuiter.git/internal/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	err := godotenv.Load()
	if err != nil {
		panic("Erro ao carregar o arquivo .env")
	}
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect postgresDB")
	}

	DB = db
}

func AutoMigrate() {
	DB.AutoMigrate(&models.User{}, &models.Hashtag{}, &models.Like{}, &models.Post{})
}
