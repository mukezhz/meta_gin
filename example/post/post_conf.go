package post

import (
	"github.com/gin-gonic/gin"
	"github.com/mukezhz/meta_gin/meta_gin"
	"gorm.io/gorm"
)

func GetPostConfig(db *gorm.DB, config *meta_gin.Config, router *gin.Engine) {
	meta_gin.SetupModelRoutes[Post, PostRequestDTO, PostResponseDTO](
		meta_gin.SetupConfig[Post, PostRequestDTO, PostResponseDTO]{
			DB:          db,
			Router:      router,
			Config:      config,
			DTOHandler:  NewPostDTOHandler(),
			Version:     "v1",
			GroupName:   "posts",
			Middlewares: []gin.HandlerFunc{meta_gin.AuthMiddleware(config, "editor")},
			Decorators:  []meta_gin.Decorator{meta_gin.NewPermissionDecorator()},
		},
	)
}
