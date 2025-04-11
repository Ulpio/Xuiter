package controllers

import (
	"github.com/Ulpio/xuiter.git/internal/services"
	"github.com/gin-gonic/gin"
)

func RegisterUser(c *gin.Context) {
	type UserSignUp struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var userSignUp UserSignUp
	if err := c.ShouldBindJSON(&userSignUp); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	id, err := services.RegisterUser(userSignUp.Username, userSignUp.Email, userSignUp.Password)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to register user"})
		return
	}
	c.JSON(200, gin.H{"message": "User registered successfully", "id": id})
}
