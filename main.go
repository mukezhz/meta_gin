package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	config, err := LoadConfig("config.toml")
	if err != nil {
		panic("failed to load config: " + err.Error())
	}

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&User{})

	router := gin.Default()
	router.Use(AuthMiddleware(config, "admin"))

	userDTOHandler := NewUserDTOHandler()
	userHandler := NewCRUDHandler[User, UserRequestDTO, UserResponseDTO](db, userDTOHandler)
	router.POST("/users", CheckPermissionDecorator(userHandler.Create()))
	router.GET("/users", AuthMiddleware(config, "editor"), CheckPermissionDecorator(userHandler.List()))

	router.Run(":8888")
}
