package post

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Name  string
	Email string
}

func (u Post) TableName() string {
	return "posts"
}
