package post

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Name        string
	Description string
	Price       float64
	Category    string
}

func (u Post) TableName() string {
	return "posts"
}
