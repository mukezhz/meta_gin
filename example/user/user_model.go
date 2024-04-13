package user

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name  string
	Email string
}

func (u User) TableName() string {
	return "users"
}
