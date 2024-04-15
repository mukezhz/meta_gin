package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mukezhz/meta_gin/example/post"
	"github.com/mukezhz/meta_gin/example/user"
	"github.com/mukezhz/meta_gin/meta_gin"
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

	m := meta_gin.NewMetaGin(router, db, config)

	user.SetupUser(m)
	post.SetupPost(m)

	m.Run(":8888")
}
