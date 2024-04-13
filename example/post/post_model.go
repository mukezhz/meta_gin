package post

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title  string
	Author string
}

func (u Post) TableName() string {
	return "posts"
}
