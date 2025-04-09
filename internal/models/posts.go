package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Content  string
	ImageURL string
	UserID   uint
	User     User
	Likes    []Like
	Hashtags []Hashtag `gorm:"many2many:post_hashtags"`
}

type Like struct {
	gorm.Model
	UserID uint
	PostID uint
}

type Hashtag struct {
	gorm.Model
	Name  string `gorm:"unique"`
	Posts []Post `gorm:"many2many:post_hashtags"`
}
