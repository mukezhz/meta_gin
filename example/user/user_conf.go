package user

import (
	"github.com/gin-gonic/gin"
	"github.com/mukezhz/meta_gin/meta_gin"
	"gorm.io/gorm"
)

func GetUserConfig(db *gorm.DB, config *meta_gin.Config, router *gin.Engine) {
	meta_gin.SetupModelRoutes[User, UserRequestDTO, UserResponseDTO](
		meta_gin.SetupConfig[User, UserRequestDTO, UserResponseDTO]{
			DB:          db,
			Router:      router,
			Config:      config,
			DTOHandler:  NewUserDTOHandler(),
			Version:     "v1",
			GroupName:   "users",
			Middlewares: []gin.HandlerFunc{meta_gin.AuthMiddleware(config, "editor")},
			Decorators:  []meta_gin.Decorator{meta_gin.NewPermissionDecorator()},
		},
	)
	// Add a custom route
	router.Group("/api").Group("/v1").Group("/users").GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello, World!"})
	})
}
