package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mukezhz/meta_gin/meta_gin"
	"github.com/mukezhz/meta_gin/post"
	"github.com/mukezhz/meta_gin/user"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	config, err := meta_gin.LoadConfig("config.toml")
	if err != nil {
		panic("failed to load config: " + err.Error())
	}

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	err = db.AutoMigrate(&user.User{}, &post.Post{})
	if err != nil {
		panic("failed to migrate database")
	}

	router := gin.Default()
	router.Use(meta_gin.AuthMiddleware(config, "admin"))

	user.GetUserConfig(db, config, router)
	post.GetPostConfig(db, config, router)

	err = router.Run(":8888")
	if err != nil {
		panic("failed to run server")
	}
}
