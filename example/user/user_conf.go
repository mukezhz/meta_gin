package user

import (
	"github.com/gin-gonic/gin"
	"github.com/mukezhz/meta_gin/meta_gin"
	"gorm.io/gorm"
)

func GetUserConfig(db *gorm.DB, config *meta_gin.Config, router *gin.Engine) {
	repository := meta_gin.NewRepository[User](db)
	service := meta_gin.NewService[User](repository)
	_ = service
	// Add a custom route
	router.Group("/api").Group("/v1").Group("/users").GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello, World!"})
	})
}
