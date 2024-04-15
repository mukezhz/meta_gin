package user

import (
	"github.com/gin-gonic/gin"
	"github.com/mukezhz/meta_gin"
)

func SetupUser(m *meta_gin.MetaGin) {
	repository := meta_gin.NewRepository[User](m.DB)
	service := meta_gin.NewService[User](repository)
	_ = service
	// Add a custom route
	m.Router.Group("/api").Group("/v1").Group("/users").GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello, World!"})
	})
}
