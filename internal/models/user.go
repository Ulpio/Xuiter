package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username  string `gorm:"unique"`
	Email     string `gorm:"unique"`
	Password  string
	Bio       string
	AvatarURL string
	Posts     []Post
	Followers []*User `gorm:"many2many:user_followers"`
}
