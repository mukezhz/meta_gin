package post

import (
	"github.com/gin-gonic/gin"
	"github.com/mukezhz/meta_gin/meta_gin"
	"gorm.io/gorm"
)

func GetPostConfig(db *gorm.DB, config *meta_gin.Config, router *gin.Engine) {
	repository := meta_gin.NewRepository[Post](db)
	service := meta_gin.NewService[Post](repository)
	PostDTOHandler := NewPostDTOHandler()
	postHandler := meta_gin.NewCRUDHandler[Post, PostRequestDTO, PostResponseDTO](db, service, PostDTOHandler)
	PostRouter := meta_gin.NewRouteHandler[Post, PostRequestDTO, PostResponseDTO](postHandler, router)
	PostConfig := meta_gin.RouteConfig{
		Version:   "v1",
		GroupName: "posts",
		Routes: []meta_gin.RouteInfo{
			{
				Path:        "/",
				Method:      "POST",
				Handler:     meta_gin.CheckPermissionDecorator(postHandler.List()),
				Middlewares: []gin.HandlerFunc{meta_gin.AuthMiddleware(config, "editor")},
			},
			{
				Path:        "/",
				Method:      "GET",
				Handler:     meta_gin.CheckPermissionDecorator(postHandler.List()),
				Middlewares: []gin.HandlerFunc{meta_gin.AuthMiddleware(config, "editor")},
			},
		},
	}
	meta_gin.RegisterRoutes[Post, PostRequestDTO, PostResponseDTO](PostRouter, PostConfig)
}
